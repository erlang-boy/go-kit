// Copyright 2013 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package params

import (
	"fmt"
	"strings"
	"time"

	"github.com/automation/errors"
	_ "github.com/automation/version"
	//	"tools"
)

type GetFarmUuidResults struct {
	FarmUuid string `json:"farm_uuid"`
}

// FindTags wraps a slice of strings that are prefixes to use when
// searching for matching tags.
type FindTags struct {
	Prefixes []string `json:"prefixes"`
}

// FindTagsResults wraps the mapping between the requested prefix and the
// matching tags for each requested prefix.
type FindTagsResults struct {
	Matches map[string][]Entity `json:"matches"`
}

// Entity identifies a single entity.
type Entity struct {
	Tag string `json:"tag"`
}

// Entities identifies multiple entities.
type Entities struct {
	Entities []Entity `json:"entities"`
}

// EntitiesResults contains multiple Entities results (where each
// Entities is the result of a query).
type EntitiesResults struct {
	Results []EntitiesResult `json:"results"`
}

// EntitiesResult is the result of one query that either yields some
// set of entities or an error.
type EntitiesResult struct {
	Entities []Entity `json:"entities"`
	Error    *Error   `json:"error,omitempty"`
}

// ErrorResults holds the results of calling a bulk operation which
// returns no data, only an error result. The order and
// number of elements matches the operations specified in the request.
type ErrorResults struct {
	// Results contains the error results from each operation.
	Results []ErrorResult `json:"results"`
}

// OneError returns the error from the result
// of a bulk operation on a single value.
func (result ErrorResults) OneError() error {
	if n := len(result.Results); n != 1 {
		return fmt.Errorf("expected 1 result, got %d", n)
	}
	if err := result.Results[0].Error; err != nil {
		return err
	}
	return nil
}

// Combine returns one error from the result which is an accumulation of the
// errors. If there are no errors in the result, the return value is nil.
// Otherwise the error values are combined with new-line characters.
func (result ErrorResults) Combine() error {
	var errorStrings []string
	for _, r := range result.Results {
		if r.Error != nil {
			errorStrings = append(errorStrings, r.Error.Error())
		}
	}
	if errorStrings != nil {
		return errors.New(strings.Join(errorStrings, "\n"))
	}
	return nil
}

// ErrorResult holds the error status of a single operation.
type ErrorResult struct {
	Error *Error `json:"error,omitempty"`
}

// PublicAddress holds parameters for the PublicAddress call.
type PublicAddress struct {
	Target string `json:"target"`
}

// PublicAddressResults holds results of the PublicAddress call.
type PublicAddressResults struct {
	PublicAddress string `json:"public-address"`
}

// PrivateAddress holds parameters for the PrivateAddress call.
type PrivateAddress struct {
	Target string `json:"target"`
}

// PrivateAddressResults holds results of the PrivateAddress call.
type PrivateAddressResults struct {
	PrivateAddress string `json:"private-address"`
}

// UpdateBehavior contains settings that are duplicated in several
// places. Let's just embed this instead.
type UpdateBehavior struct {
	EnableOSRefreshUpdate bool `json:"enable-os-refresh-update"`
	EnableOSUpgrade       bool `json:"enable-os-upgrade"`
}

// FacadeVersions describes the available Facades and what versions of each one
// are available
type FacadeVersions struct {
	Name     string `json:"name"`
	Versions []int  `json:"versions"`
}

// LogRecord is used to transmit log messages to the logsink API
// endpoint.  Single character field names are used for serialisation
// to keep the size down. These messages are going to be sent a lot.
type LogRecord struct {
	Time     time.Time `json:"t"`
	Module   string    `json:"m"`
	Location string    `json:"l"`
	Level    string    `json:"v"`
	Message  string    `json:"x"`
	Entity   string    `json:"e,omitempty"`
}

type Presence struct {
	Result string `json:"result"`
}
