// Copyright 2014 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package util

import (
	"fmt"
	//	"io"
	//	"strconv"

	"github.com/juju/errors"
	"github.com/juju/loggo"

	//	"agent"
	"go-kit/actor"
)

var logger = loggo.GetLogger("automation.cmd.automationd.util")

// RequiredError is useful when complaining about missing command-line options.
func RequiredError(name string) error {
	return fmt.Errorf("--%s option must be set", name)
}

// IsFatal determines if an error is fatal to the process.
func IsFatal(err error) bool {
	err = errors.Cause(err)
	switch err {
	case actor.ErrTerminateAgent:
		return true
	}

	_, ok := err.(*FatalError)
	return ok
}

// FatalError is an error type designated for fatal errors.
type FatalError struct {
	Err string
}

// Error returns an error string.
func (e *FatalError) Error() string {
	return e.Err
}

func importance(err error) int {
	err = errors.Cause(err)
	switch {
	case err == nil:
		return 0
	default:
		return 1
	case err == actor.ErrTerminateAgent:
		return 4
	}
}

// MoreImportant returns whether err0 is more important than err1 -
// that is, whether we should act on err0 in preference to err1.
func MoreImportant(err0, err1 error) bool {
	return importance(err0) > importance(err1)
}

// MoreImportantError returns the most important error
func MoreImportantError(err0, err1 error) error {
	if importance(err0) > importance(err1) {
		return err0
	}
	return err1
}

func DaemonDone(logger loggo.Logger, err error) error {
	err = errors.Cause(err)
	switch err {
	case actor.ErrTerminateAgent:
		// These errors are swallowed here because we want to exit
		// the agent process without error, to avoid the init system
		// restarting us.
		err = nil
	}
	return err
}

// Breakable provides a type that exposes an IsBroken check.
type Breakable interface {
	IsBroken() bool
}

func ConnectionIsFatal(logger loggo.Logger, conns ...Breakable) func(err error) bool {
	return func(err error) bool {
		if IsFatal(err) {
			return true
		}
		for _, conn := range conns {
			if ConnectionIsDead(logger, conn) {
				return true
			}
		}
		return false
	}
}

// ConnectionIsDead returns true if the given Breakable is broken.
var ConnectionIsDead = func(logger loggo.Logger, conn Breakable) bool {
	return conn.IsBroken()
}

// Pinger provides a type that knows how to ping.
type Pinger interface {
	Ping() error
}

// TODO(mjs) - this only exists for checking State instances in the
// machine agent and won't be necessary once either:
// 1. State grows a Broken() channel like api.Connection (which is
//    actually quite a nice idea).
// 2. The dependency engine conversion is completed for the machine
//    agent.
func PingerIsFatal(logger loggo.Logger, conns ...Pinger) func(err error) bool {
	return func(err error) bool {
		if IsFatal(err) {
			return true
		}
		for _, conn := range conns {
			if PingerIsDead(logger, conn) {
				return true
			}
		}
		return false
	}
}

// PingerIsDead returns true if the given pinger fails to ping.
var PingerIsDead = func(logger loggo.Logger, conn Pinger) bool {
	if err := conn.Ping(); err != nil {
		logger.Infof("error pinging %T: %v", conn, err)
		return true
	}
	return false
}
