// Copyright 2012-2015 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package api

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"sync/atomic"
	"time"

	"github.com/automation/errors"
	"github.com/automation/loggo"
	"github.com/automation/retry"
	"github.com/automation/utils"
	"github.com/automation/utils/clock"
	"github.com/automation/utils/parallel"
	"github.com/automation/version"
	"golang.org/x/net/websocket"
	"gopkg.in/automation/names.v2"

	"api/base"
	"automation/common/network"
	//"apiserver/observer"
	"api/params"
	"automation/common/rpc"
	"automation/common/rpc/jsoncodec"
)

// PingPeriod defines how often the internal connection health check
// will run.
//const PingPeriod = 1 * time.Minute
const PingPeriod = 15 * time.Second

// pingTimeout defines how long a health check can take before we
// consider it to have failed.
const pingTimeout = 30 * time.Second

var logger = loggo.GetLogger("automation.api")

type rpcConnection interface {
	Call(req rpc.Request, params, response interface{}) error
	Dead() <-chan struct{}
	Close() error
}

// state is the internal implementation of the Connection interface.
type state struct {
	client rpcConnection
	conn   *websocket.Conn
	clock  clock.Clock

	// addr is the address used to connect to the API server.
	addr string

	// serverVersion holds the version of the API server that we are
	// connected to.  It is possible that this version is 0 if the
	// server does not report this during login.
	serverVersion version.Number

	// hostPorts is the API server addresses returned from Login,
	// which the client may cache and use for failover.
	hostPorts [][]network.HostPort

	// facadeVersions holds the versions of all facades as reported by
	// Login
	facadeVersions map[string][]int

	// pingFacadeVersion is the version to use for the pinger. This is lazily
	// set at initialization to avoid a race in our tests. See
	// http://pad.lv/1614732 for more details regarding the race.
	pingerFacadeVersion int

	// broken is a channel that gets closed when the connection is
	// broken.
	broken chan struct{}

	// closed is a channel that gets closed when State.Close is called.
	closed chan struct{}

	// loggedIn holds whether the client has successfully logged
	// in. It's a int32 so that the atomic package can be used to
	// access it safely.
	loggedIn int32

	// tag, password, macaroons and nonce hold the cached login
	// credentials. These are only valid if loggedIn is 1.
	tag   string
	nonce string

	// serverRootAddress holds the cached API server address and port used
	// to login.
	serverRootAddress string

	// serverScheme is the URI scheme of the API Server
	serverScheme string
}

// RedirectError is returned from Open when the controller
// needs to inform the client that the model is hosted
// on a different set of API addresses.
type RedirectError struct {
	// Servers holds the sets of addresses of the redirected
	// servers.
	Servers [][]network.HostPort
}

func (e *RedirectError) Error() string {
	return fmt.Sprintf("redirection to alternative server required")
}

// Open establishes a connection to the API server using the Info
// given, returning a State instance which can be used to make API
// requests.
//
// If the model is hosted on a different server, Open
// will return an error with a *RedirectError cause
// holding the details of another server to connect to.
//
// See Connect for details of the connection mechanics.
func Open(info *Info, opts DialOpts) (Connection, error) {
	return open(info, opts, clock.WallClock)
}

// open is the unexported version of open that also includes
// an explicit clock instance argument.
func open(
	info *Info,
	opts DialOpts,
	clock clock.Clock,
) (Connection, error) {
	if err := info.Validate(); err != nil {
		return nil, errors.Annotate(err, "validating info for opening an API connection")
	}
	if clock == nil {
		return nil, errors.NotValidf("nil clock")
	}
	conn, err := dialAPI(info, opts)
	if err != nil {
		return nil, errors.Trace(err)
	}

	//client := rpc.NewConn(jsoncodec.NewWebsocket(conn), observer.None())
	client := rpc.NewConn(jsoncodec.NewWebsocket(conn), nil)
	client.Start()

	apiHost := conn.Config().Location.Host
	st := &state{
		client:              client,
		conn:                conn,
		clock:               clock,
		addr:                apiHost,
		pingerFacadeVersion: facadeVersions["Pinger"],
		serverScheme:        "http",
		serverRootAddress:   conn.Config().Location.Host,
		tag:                 tagToString(info.Tag),
		nonce:               info.Nonce,
	}

	st.broken = make(chan struct{})
	st.closed = make(chan struct{})

	go (&monitor{
		clock:       clock,
		ping:        st.Ping,
		pingPeriod:  PingPeriod,
		pingTimeout: pingTimeout,
		closed:      st.closed,
		dead:        client.Dead(),
		broken:      st.broken,
	}).run()
	return st, nil
}

// hostSwitchingTransport provides an http.RoundTripper
// that chooses an actual RoundTripper to use
// depending on the destination host.
//
// This makes it possible to use a different set of root
// CAs for the API and all other hosts.
type hostSwitchingTransport struct {
	primaryHost string
	primary     http.RoundTripper
	fallback    http.RoundTripper
}

// RoundTrip implements http.RoundTripper.RoundTrip.
func (t *hostSwitchingTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	if req.URL.Host == t.primaryHost {
		return t.primary.RoundTrip(req)
	}
	return t.fallback.RoundTrip(req)
}

// ConnectStream implements StreamConnector.ConnectStream.
func (st *state) ConnectStream(path string, attrs url.Values) (base.Stream, error) {
	conn, err := st.connectStreamWithRetry(path, attrs, nil)
	if err != nil {
		return nil, errors.Trace(err)
	}
	return conn, nil
}

func (st *state) connectStreamWithRetry(path string, attrs url.Values, headers http.Header) (base.Stream, error) {
	if !st.isLoggedIn() {
		return nil, errors.New("cannot use ConnectStream without logging in")
	}
	// We use the standard "macaraq" macaroon authentication dance here.
	// That is, we attach any macaroons we have to the initial request,
	// and if that succeeds, all's good. If it fails with a DischargeRequired
	// error, the response will contain a macaroon that, when discharged,
	// may allow access, so we discharge it (using bakery.Client.HandleError)
	// and try the request again.
	conn, err := st.connectStream(path, attrs, headers)
	if err == nil {
		return conn, err
	}
	//if err := st.bakeryClient.HandleError(st.cookieURL, bakeryError(err)); err != nil {
	//	return nil, errors.Trace(err)
	//}
	// Try again with the discharged macaroon.
	conn, err = st.connectStream(path, attrs, headers)
	if err != nil {
		return nil, errors.Trace(err)
	}
	return conn, nil
}

// connectStream is the internal version of ConnectStream. It differs from
// ConnectStream only in that it will not retry the connection if it encounters
// discharge-required error.
func (st *state) connectStream(path string, attrs url.Values, extraHeaders http.Header) (base.Stream, error) {
	target := url.URL{
		Scheme:   "wss",
		Host:     st.addr,
		Path:     path,
		RawQuery: attrs.Encode(),
	}
	// TODO(macgreagoir) IPv6. Ubuntu still always provides IPv4 loopback,
	// and when/if this changes localhost should resolve to IPv6 loopback
	// in any case (lp:1644009). Review.
	cfg, err := websocket.NewConfig(target.String(), "http://localhost/")
	// Add any cookies because they will not be sent to websocket
	// connections by default.
	for header, values := range extraHeaders {
		for _, value := range values {
			cfg.Header.Add(header, value)
		}
	}

	connection, err := websocketDialConfig(cfg)
	if err != nil {
		return nil, err
	}
	if err := readInitialStreamError(connection); err != nil {
		return nil, errors.Trace(err)
	}
	return connection, nil
}

// readInitialStreamError reads the initial error response
// from a stream connection and returns it.
func readInitialStreamError(conn io.Reader) error {
	// We can use bufio here because the websocket guarantees that a
	// single read will not read more than a single frame; there is
	// no guarantee that a single read might not read less than the
	// whole frame though, so using a single Read call is not
	// correct. By using ReadSlice rather than ReadBytes, we
	// guarantee that the error can't be too big (>4096 bytes).
	line, err := bufio.NewReader(conn).ReadSlice('\n')
	if err != nil {
		return errors.Annotate(err, "unable to read initial response")
	}
	var errResult params.ErrorResult
	if err := json.Unmarshal(line, &errResult); err != nil {
		return errors.Annotate(err, "unable to unmarshal initial response")
	}
	if errResult.Error != nil {
		return errResult.Error
	}
	return nil
}

// apiEndpoint returns a URL that refers to the given API slash-prefixed
// endpoint path and query parameters.
func (st *state) apiEndpoint(path, query string) (*url.URL, error) {
	return &url.URL{
		Scheme:   st.serverScheme,
		Host:     st.Addr(),
		Path:     path,
		RawQuery: query,
	}, nil
}

// Ping implements api.Connection.
func (s *state) Ping() error {
	return s.APICall("Pinger", s.pingerFacadeVersion, "", "Ping", nil, nil)
}

// tagToString returns the value of a tag's String method, or "" if the tag is nil.
func tagToString(tag names.Tag) string {
	if tag == nil {
		return ""
	}
	return tag.String()
}

// dialAPI establishes a websocket connection to the RPC
// API websocket on the API server using Info. If multiple API addresses
// are provided in Info they will be tried concurrently - the first successful
// connection wins.
//
// It also returns the TLS configuration that it has derived from the Info.
func dialAPI(info *Info, opts DialOpts) (*websocket.Conn, error) {
	// Set opts.DialWebsocket here rather than in open because
	// some tests call dialAPI directly.
	if opts.DialWebsocket == nil {
		opts.DialWebsocket = websocket.DialConfig
	}
	if len(info.Addrs) == 0 {
		return nil, errors.New("no API addresses to connect to")
	}

	path := "/api"
	conn, err := dialWebsocketMulti(info.Addrs, path, opts)
	if err != nil {
		return nil, errors.Trace(err)
	}
	logger.Infof("connection established to %q", conn.RemoteAddr())
	return conn, nil
}

// dialWebsocketMulti dials a websocket with one of the provided addresses, the
// specified URL path, TLS configuration, and dial options. Each of the
// specified addresses will be attempted concurrently, and the first
// successful connection will be returned.
func dialWebsocketMulti(addrs []string, path string, opts DialOpts) (*websocket.Conn, error) {
	// Dial all addresses at reasonable intervals.
	try := parallel.NewTry(0, nil)
	defer try.Kill()
	for _, addr := range addrs {
		err := startDialWebsocket(try, addr, path, opts)
		if err == parallel.ErrStopped {
			break
		}
		if err != nil {
			return nil, errors.Trace(err)
		}
		select {
		case <-time.After(opts.DialAddressInterval):
		case <-try.Dead():
		}
	}
	try.Close()
	result, err := try.Result()
	if err != nil {
		return nil, errors.Trace(err)
	}
	return result.(*websocket.Conn), nil
}

// startDialWebsocket starts websocket connection to a single address
// on the given try instance.
func startDialWebsocket(try *parallel.Try, addr, path string, opts DialOpts) error {
	// origin is required by the WebSocket API, used for "origin policy"
	// in websockets. We pass localhost to satisfy the API; it is
	// inconsequential to us.
	const origin = "http://localhost/"
	//cfg, err := websocket.NewConfig("wss://"+addr+path, origin)
	cfg, err := websocket.NewConfig("ws://"+addr+path, origin)
	if err != nil {
		return errors.Trace(err)
	}
	return try.Start(newWebsocketDialer(cfg, opts))
}

// newWebsocketDialer0 returns a function that dials the websocket represented
// by the given configuration with the given dial options, suitable for passing
// to utils/parallel.Try.Start.
func newWebsocketDialer(cfg *websocket.Config, opts DialOpts) func(<-chan struct{}) (io.Closer, error) {
	// TODO(katco): 2016-08-09: lp:1611427
	openAttempt := utils.AttemptStrategy{
		Total: opts.Timeout,
		Delay: opts.RetryDelay,
	}

	if openAttempt.Min == 0 && openAttempt.Delay > 0 {
		openAttempt.Min = int(openAttempt.Total / openAttempt.Delay)
	}

	return func(stop <-chan struct{}) (io.Closer, error) {
		for a := openAttempt.Start(); a.Next(); {
			select {
			case <-stop:
				return nil, parallel.ErrStopped
			default:
			}
			logger.Infof("dialing %q", cfg.Location)
			conn, err := opts.DialWebsocket(cfg)
			if err == nil {
				return conn, nil
			}
			if !a.HasNext() {
				logger.Errorf("error dialing %q: %v", cfg.Location, err)
				return nil, errors.Annotatef(err, "unable to connect to API")
			}
		}
		panic("unreachable")
	}
}

type hasErrorCode interface {
	ErrorCode() string
}

// APICall places a call to the remote machine.
//
// This fills out the rpc.Request on the given facade, version for a given
// object id, and the specific RPC method. It marshalls the Arguments, and will
// unmarshall the result into the response object that is supplied.
func (s *state) APICall(facade string, version int, id, method string, args, response interface{}) error {
	retrySpec := retry.CallArgs{
		Func: func() error {
			return s.client.Call(rpc.Request{
				Type:    facade,
				Version: version,
				Id:      id,
				Action:  method,
			}, args, response)
		},
		IsFatalError: func(err error) bool {
			err = errors.Cause(err)
			ec, ok := err.(hasErrorCode)
			if !ok {
				return true
			}
			return ec.ErrorCode() != params.CodeRetry
		},
		Delay:       100 * time.Millisecond,
		MaxDelay:    1500 * time.Millisecond,
		MaxDuration: 10 * time.Second,
		BackoffFunc: retry.DoubleDelay,
		Clock:       s.clock,
	}
	err := retry.Call(retrySpec)
	return errors.Trace(err)
}

func (s *state) Close() error {
	err := s.client.Close()
	select {
	case <-s.closed:
	default:
		close(s.closed)
	}
	<-s.broken
	return err
}

// Broken implements api.Connection.
func (s *state) Broken() <-chan struct{} {
	return s.broken
}

// IsBroken implements api.Connection.
func (s *state) IsBroken() bool {
	select {
	case <-s.broken:
		return true
	default:
	}
	if err := s.Ping(); err != nil {
		logger.Debugf("connection ping failed: %v", err)
		return true
	}
	return false
}

// Addr returns the address used to connect to the API server.
func (s *state) Addr() string {
	return s.addr
}

// APIHostPorts returns addresses that may be used to connect
// to the API server, including the address used to connect.
//
// The addresses are scoped (public, cloud-internal, etc.), so
// the client may choose which addresses to attempt. For the
// Automation CLI, all addresses must be attempted, as the CLI may
// be invoked both within and outside the model (think
// private clouds).
func (s *state) APIHostPorts() [][]network.HostPort {
	// NOTE: We're making a copy of s.hostPorts before returning it,
	// for safety.
	hostPorts := make([][]network.HostPort, len(s.hostPorts))
	for i, server := range s.hostPorts {
		hostPorts[i] = append([]network.HostPort{}, server...)
	}
	return hostPorts
}

// AllFacadeVersions returns what versions we know about for all facades
func (s *state) AllFacadeVersions() map[string][]int {
	facades := make(map[string][]int, len(s.facadeVersions))
	for name, versions := range s.facadeVersions {
		facades[name] = append([]int{}, versions...)
	}
	return facades
}

// BestFacadeVersion compares the versions of facades that we know about, and
// the versions available from the server, and reports back what version is the
// 'best available' to use.
// TODO(jam) this is the eventual implementation of what version of a given
// Facade we will want to use. It needs to line up the versions that the server
// reports to us, with the versions that our client knows how to use.
func (s *state) BestFacadeVersion(facade string) int {
	return bestVersion(facadeVersions[facade], s.facadeVersions[facade])
}

// serverRoot returns the cached API server address and port used
// to login, prefixed with "<URI scheme>://" (usually https).
func (s *state) serverRoot() string {
	return s.serverScheme + "://" + s.serverRootAddress
}

func (s *state) isLoggedIn() bool {
	return atomic.LoadInt32(&s.loggedIn) == 1
}

func (s *state) setLoggedIn() {
	atomic.StoreInt32(&s.loggedIn, 1)
}
