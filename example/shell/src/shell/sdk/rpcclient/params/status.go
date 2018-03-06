// Copyright 2015 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package params

//// TODO(ericsnow) Eliminate the automation-related imports.
//
import (
	"time"
)

// StatusParams holds parameters for the Status call.
type StatusParams struct {
	Patterns []string `json:"patterns"`
}

// StatusResult holds an entity status, extra information, or an
// error.
type StatusResult struct {
	Error  *Error                 `json:"error,omitempty"`
	Id     string                 `json:"id"`
	Life   Life                   `json:"life"`
	Status string                 `json:"status"`
	Info   string                 `json:"info"`
	Data   map[string]interface{} `json:"data"`
	Since  *time.Time             `json:"since"`
}

// StatusResults holds multiple status results.
type StatusResults struct {
	Results []StatusResult `json:"results"`
}

// Life describes the lifecycle state of an entity ("alive", "dying" or "dead").
type Life string

const (
	Alive   Life = "alive"
	Dying   Life = "dying"
	Dead    Life = "dead"
	Halting Life = "halting"
	Halted  Life = "halted"
)
