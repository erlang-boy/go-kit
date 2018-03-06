// Copyright 2014 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package api

import (
	_ "bytes"
	_ "encoding/base64"
	"encoding/json"
	"io"
	"net/http"
	"net/url"

	"github.com/automation/errors"
	"github.com/automation/httprequest"

	"api/params"
)

// HTTPClient implements Connection.APICaller.HTTPClient and returns an HTTP
// client pointing to the API server "/model/:uuid/" path.
func (s *state) HTTPClient() (*httprequest.Client, error) {
	baseURL, err := s.apiEndpoint("/", "")
	if err != nil {
		return nil, errors.Trace(err)
	}
	return s.httpClient(baseURL)
}

// HTTPClient implements Connection.APICaller.HTTPClient and returns an HTTP
// client pointing to the API server root path.
func (s *state) RootHTTPClient() (*httprequest.Client, error) {
	return s.httpClient(&url.URL{
		Scheme: s.serverScheme,
		Host:   s.Addr(),
	})
}

func (s *state) httpClient(baseURL *url.URL) (*httprequest.Client, error) {
	if !s.isLoggedIn() {
		return nil, errors.New("no HTTP client available without logging in")
	}
	return &httprequest.Client{
		BaseURL: baseURL.String(),
		Doer: httpRequestDoer{
			st: s,
		},
		UnmarshalError: unmarshalHTTPErrorResponse,
	}, nil
}

// httpRequestDoer implements httprequest.Doer and httprequest.DoerWithBody
// by using httpbakery and the state to make authenticated requests to
// the API server.
type httpRequestDoer struct {
	st *state
}

var _ httprequest.Doer = httpRequestDoer{}

var _ httprequest.DoerWithBody = httpRequestDoer{}

// Do implements httprequest.Doer.Do.
func (doer httpRequestDoer) Do(req *http.Request) (*http.Response, error) {
	return doer.DoWithBody(req, nil)
}

// DoWithBody implements httprequest.DoerWithBody.DoWithBody.
func (doer httpRequestDoer) DoWithBody(req *http.Request, body io.ReadSeeker) (*http.Response, error) {
	//TODO: to be implemented
	return nil, nil
}

// unmarshalHTTPErrorResponse unmarshals an error response from
// an HTTP endpoint. For historical reasons, these endpoints
// return several different incompatible error response formats.
// We cope with this by accepting all of the possible formats
// and unmarshaling accordingly.
//
// It always returns a non-nil error.
func unmarshalHTTPErrorResponse(resp *http.Response) error {
	var body json.RawMessage
	if err := httprequest.UnmarshalJSONResponse(resp, &body); err != nil {
		return errors.Trace(err)
	}
	// genericErrorResponse defines a struct that is compatible with all the
	// known error types, so that we can know which of the
	// possible error types has been returned.
	//
	// Another possible approach might be to look at resp.Request.URL.Path
	// and determine the expected error type from that, but that
	// seems more fragile than this approach.
	type genericErrorResponse struct {
		Error json.RawMessage `json:"error"`
	}
	var generic genericErrorResponse
	if err := json.Unmarshal(body, &generic); err != nil {
		return errors.Annotatef(err, "incompatible error response")
	}
	var errorBody []byte
	if len(generic.Error) > 0 {
		// We have an Error field, therefore the error must be in that.
		// (it's a params.ErrorResponse)
		errorBody = generic.Error
	} else {
		// There wasn't an Error field, so the error must be directly
		// in the body of the response.
		errorBody = body
	}
	var perr params.Error
	if err := json.Unmarshal(errorBody, &perr); err != nil {
		return errors.Annotatef(err, "incompatible error response")
	}
	if perr.Message == "" {
		return errors.Errorf("error response with no message")
	}
	return &perr
}

// HTTPDoer exposes the functionality of httprequest.Client needed here.
type HTTPDoer interface {
	// Do sends the given request.
	Do(req *http.Request, body io.ReadSeeker, resp interface{}) error
}

// openBlob streams the identified blob from the controller via the
// provided HTTP client.
func openBlob(httpClient HTTPDoer, endpoint string, args url.Values) (io.ReadCloser, error) {
	apiURL, err := url.Parse(endpoint)
	if err != nil {
		return nil, errors.Trace(err)
	}
	apiURL.RawQuery = args.Encode()
	req, err := http.NewRequest("GET", apiURL.String(), nil)
	if err != nil {
		return nil, errors.Annotate(err, "cannot create HTTP request")
	}

	var resp *http.Response
	if err := httpClient.Do(req, nil, &resp); err != nil {
		return nil, errors.Trace(err)
	}
	return resp.Body, nil
}