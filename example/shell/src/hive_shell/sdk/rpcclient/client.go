// Copyright 2013, 2014 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package api

import (
	"github.com/automation/errors"
	"github.com/automation/version"
	"golang.org/x/net/websocket"

	"api/base"
	"api/params"
)

// Client represents the client-accessible part of the state.
type Client struct {
	base.ClientFacade
	facade base.FacadeCaller
	st     *state
}

// Status returns the status of the automation model.
func (c *Client) Status(patterns []string) (*params.StatusResults, error) {
	var result params.StatusResults
	p := params.StatusParams{Patterns: patterns}
	if err := c.facade.FacadeCall("Status", p, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Close closes the Client's underlying State connection
// Client is unique among the api.State facades in closing its own State
// connection, but it is conventional to use a Client object without any access
// to its underlying state connection.
func (c *Client) Close() error {
	return c.st.Close()
}

type minAutomationVersionErr struct {
	*errors.Err
}

// AgentVersion reports the version number of the api server.
func (c *Client) AgentVersion() (version.Number, error) {
	var result params.AgentVersionResult
	if err := c.facade.FacadeCall("AgentVersion", nil, &result); err != nil {
		return version.Number{}, err
	}
	return result.Version, nil
}

// websocketDialConfig is called instead of websocket.DialConfig so we can
// override it in tests.
var websocketDialConfig = func(config *websocket.Config) (base.Stream, error) {
	c, err := websocket.DialConfig(config)
	if err != nil {
		return nil, errors.Trace(err)
	}
	return websocketStream{c}, nil
}

type websocketStream struct {
	*websocket.Conn
}

func (c websocketStream) ReadJSON(v interface{}) error {
	return websocket.JSON.Receive(c.Conn, v)
}

func (c websocketStream) WriteJSON(v interface{}) error {
	return websocket.JSON.Send(c.Conn, v)
}
