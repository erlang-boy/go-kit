// Copyright 2015 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package base

import (
	"status"
	"time"
)

// Status represents the status of a machine, application, or unit.
type Status struct {
	Status status.Status
	Info   string
	Data   map[string]interface{}
	Since  *time.Time
}
