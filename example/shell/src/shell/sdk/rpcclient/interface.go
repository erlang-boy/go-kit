// Copyright 2012-2015 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package api

import (
	_ "net/url"
	"time"

	"github.com/automation/errors"
	"github.com/automation/version"
	"golang.org/x/net/websocket"
	"gopkg.in/automation/names.v2"

	"api/base"
	"automation/common/network"
	//	"api/discoverspaces"
	//	"api/reboot"
	//	"api/unitassigner"
	//	"api/uniter"
	"github.com/automation/utils/set"
)

// Info encapsulates information about a server holding automation state and
// can be used to make a connection to it.
type Info struct {

	// This block of fields is sufficient to connect:

	// Addrs holds the addresses of the controllers.
	Addrs []string

	// Tag holds the name of the entity that is connecting.
	// If this is nil, and the password is empty, macaroon authentication
	// will be used to log in unless SkipLogin is true.
	Tag names.Tag

	// Nonce holds the nonce used when provisioning the machine. Used
	// only by the machine agent.
	Nonce string `yaml:",omitempty"`
}

// Ports returns the unique ports for the api addresses.
func (info *Info) Ports() []int {
	ports := set.NewInts()
	hostPorts, err := network.ParseHostPorts(info.Addrs...)
	if err != nil {
		// Addresses have already been validated.
		panic(err)
	}
	for _, hp := range hostPorts {
		ports.Add(hp.Port)
	}
	return ports.Values()
}

// Validate validates the API info.
func (info *Info) Validate() error {
	if len(info.Addrs) == 0 {
		return errors.NotValidf("missing addresses")
	}
	if _, err := network.ParseHostPorts(info.Addrs...); err != nil {
		return errors.NotValidf("host addresses: %v", err)
	}
	return nil
}

// DialOpts holds configuration parameters that control the
// Dialing behavior when connecting to a controller.
type DialOpts struct {
	// DialAddressInterval is the amount of time to wait
	// before starting to dial another address.
	DialAddressInterval time.Duration

	// Timeout is the amount of time to wait contacting
	// a controller.
	Timeout time.Duration

	// RetryDelay is the amount of time to wait between
	// unsuccessful connection attempts.
	RetryDelay time.Duration

	// DialWebsocket is used to make connections to API servers.
	// It will be called with a websocket URL to connect to,
	// and the TLS configuration to use to secure the connection.
	//
	// If DialWebsocket is nil, webaocket.DialConfig will be used.
	//
	// This field is provided for testing purposes only.
	DialWebsocket func(cfg *websocket.Config) (*websocket.Conn, error)
}

// DefaultDialOpts returns a DialOpts representing the default
// parameters for contacting a controller.
func DefaultDialOpts() DialOpts {
	return DialOpts{
		DialAddressInterval: 50 * time.Millisecond,
		Timeout:             10 * time.Minute,
		RetryDelay:          2 * time.Second,
	}
}

// OpenFunc is the usual form of a function that opens an API connection.
type OpenFunc func(*Info, DialOpts) (Connection, error)

// Connection exists purely to make api-opening funcs mockable. It's just a
// dumb copy of all the methods on api.state; we can and should be extracting
// smaller and more relevant interfaces (and dropping some of them too).

// Connection represents a connection to a Automation API server.
type Connection interface {

	// This first block of methods is pretty close to a sane Connection interface.
	Close() error
	Addr() string
	APIHostPorts() [][]network.HostPort

	// Broken returns a channel which will be closed if the connection
	// is detected to be broken, either because the underlying
	// connection has closed or because API pings have failed.
	Broken() <-chan struct{}

	// IsBroken returns whether the connection is broken. It checks
	// the Broken channel and if that is open, attempts a connection
	// ping.
	IsBroken() bool

	ServerVersion() (version.Number, bool)

	// APICaller provides the facility to make API calls directly.
	// This should not be used outside the api/* packages or tests.
	base.APICaller

	// All the rest are strange and questionable and deserve extra attention
	// and/or discussion.

	// Ping makes an API request which checks if the connection is
	// still functioning.
	// NOTE: This method is deprecated. Please use IsBroken or Broken instead.
	Ping() error

	// I think this is actually dead code. It's tested, at least, so I'm
	// keeping it for now, but it's not apparently used anywhere else.
	AllFacadeVersions() map[string][]int

	Client() *Client
	//Uniter() (*uniter.State, error)
	//Reboot() (reboot.State, error)
	//DiscoverSpaces() *discoverspaces.API
	//UnitAssigner() unitassigner.API
}
