// Copyright 2013 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package api

import (
	"net"
	_ "net/url"
	"strconv"

	_ "github.com/automation/errors"
	"github.com/automation/version"
	_ "gopkg.in/automation/names.v2"

	"api/base"
	"automation/common/network"
	//	"api/discoverspaces"
	//	"api/imagemetadata"
	//	"api/instancepoller"
	//	"api/keyupdater"
	//	"api/reboot"
	//	"api/unitassigner"
	//	"api/uniter"
	_ "api/params"
)

// slideAddressToFront moves the address at the location (serverIndex, addrIndex) to be
// the first address of the first server.
func slideAddressToFront(servers [][]network.HostPort, serverIndex, addrIndex int) {
	server := servers[serverIndex]
	hostPort := server[addrIndex]
	// Move the matching address to be the first in this server
	for ; addrIndex > 0; addrIndex-- {
		server[addrIndex] = server[addrIndex-1]
	}
	server[0] = hostPort
	for ; serverIndex > 0; serverIndex-- {
		servers[serverIndex] = servers[serverIndex-1]
	}
	servers[0] = server
}

// addAddress appends a new server derived from the given
// address to servers if the address is not already found
// there.
func addAddress(servers [][]network.HostPort, addr string) ([][]network.HostPort, error) {
	for i, server := range servers {
		for j, hostPort := range server {
			if hostPort.NetAddr() == addr {
				slideAddressToFront(servers, i, j)
				return servers, nil
			}
		}
	}
	host, portString, err := net.SplitHostPort(addr)
	if err != nil {
		return nil, err
	}
	port, err := strconv.Atoi(portString)
	if err != nil {
		return nil, err
	}
	result := make([][]network.HostPort, 0, len(servers)+1)
	result = append(result, network.NewHostPorts(port, host))
	result = append(result, servers...)
	return result, nil
}

// Client returns an object that can be used
// to access client-specific functionality.
func (st *state) Client() *Client {
	frontend, backend := base.NewClientFacade(st, "Client")
	return &Client{ClientFacade: frontend, facade: backend, st: st}
}

// ServerVersion holds the version of the API server that we are connected to.
// It is possible that this version is Zero if the server does not report this
// during login. The second result argument indicates if the version number is
// set.
func (st *state) ServerVersion() (version.Number, bool) {
	return st.serverVersion, st.serverVersion != version.Zero
}
