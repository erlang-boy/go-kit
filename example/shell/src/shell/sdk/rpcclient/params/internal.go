// Copyright 2013 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package params

import (
	"time"

	"github.com/automation/version"

	"status"
)

// StringsResult holds the result of an API call that returns a slice
// of strings or an error.
type StringsResult struct {
	Error  *Error   `json:"error,omitempty"`
	Result []string `json:"result,omitempty"`
}

// StringsResults holds the bulk operation result of an API call
// that returns a slice of strings or an error.
type StringsResults struct {
	Results []StringsResult `json:"results"`
}

// StringResult holds a string or an error.
type StringResult struct {
	Error  *Error `json:"error,omitempty"`
	Result string `json:"result"`
}

// StringResults holds the bulk operation result of an API call
// that returns a string or an error.
type StringResults struct {
	Results []StringResult `json:"results"`
}

// MapResult holds a generic map or an error.
type MapResult struct {
	Result map[string]interface{} `json:"result"`
	Error  *Error                 `json:"error,omitempty"`
}

// MapResults holds the bulk operation result of an API call
// that returns a map or an error.
type MapResults struct {
	Results []MapResult `json:"results"`
}

// StringBoolResult holds the result of an API call that returns a
// string and a boolean.
type StringBoolResult struct {
	Error  *Error `json:"error,omitempty"`
	Result string `json:"result"`
	Ok     bool   `json:"ok"`
}

// StringBoolResults holds multiple results with a string and a bool
// each.
type StringBoolResults struct {
	Results []StringBoolResult `json:"results"`
}

// BoolResult holds the result of an API call that returns a
// a boolean or an error.
type BoolResult struct {
	Error  *Error `json:"error,omitempty"`
	Result bool   `json:"result"`
}

// BoolResults holds multiple results with BoolResult each.
type BoolResults struct {
	Results []BoolResult `json:"results"`
}

// IntResults holds multiple results with an int in each.
type IntResults struct {
	// Results holds a list of results for calls that return an int or error.
	Results []IntResult `json:"results"`
}

// IntResult holds the result of an API call that returns a
// int or an error.
type IntResult struct {
	// Error holds the error (if any) of this call.
	Error *Error `json:"error,omitempty"`
	// Result holds the integer result of the call (if Error is nil).
	Result int `json:"result"`
}

// Settings holds relation settings names and values.
type Settings map[string]string

// SettingsResult holds a relation settings map or an error.
type SettingsResult struct {
	Error    *Error   `json:"error,omitempty"`
	Settings Settings `json:"settings"`
}

// SettingsResults holds the result of an API calls that
// returns settings for multiple relations.
type SettingsResults struct {
	Results []SettingsResult `json:"results"`
}

// ConfigSettings holds unit, application or cham configuration settings
// with string keys and arbitrary values.
type ConfigSettings map[string]interface{}

// ConfigSettingsResult holds a configuration map or an error.
type ConfigSettingsResult struct {
	Error    *Error         `json:"error,omitempty"`
	Settings ConfigSettings `json:"settings"`
}

// ConfigSettingsResults holds multiple configuration maps or errors.
type ConfigSettingsResults struct {
	Results []ConfigSettingsResult `json:"results"`
}

// BytesResult holds the result of an API call that returns a slice
// of bytes.
type BytesResult struct {
	Result []byte `json:"result"`
}

// LifeResult holds the life status of a single entity, or an error
// indicating why it is not available.
type LifeResult struct {
	Life  Life   `json:"life"`
	Error *Error `json:"error,omitempty"`
}

// LifeResults holds the life or error status of multiple entities.
type LifeResults struct {
	Results []LifeResult `json:"results"`
}

// EntityStatus holds the status of an entity.
type EntityStatus struct {
	Status status.Status          `json:"status"`
	Info   string                 `json:"info"`
	Data   map[string]interface{} `json:"data,omitempty"`
	Since  *time.Time             `json:"since"`
}

// EntityStatusArgs holds parameters for setting the status of a single entity.
type EntityStatusArgs struct {
	Tag    string                 `json:"tag"`
	Status string                 `json:"status"`
	Info   string                 `json:"info"`
	Data   map[string]interface{} `json:"data"`
}

// SetStatus holds the parameters for making a SetStatus/UpdateStatus call.
type SetStatus struct {
	Entities []EntityStatusArgs `json:"entities"`
}

// AgentGetEntitiesResults holds the results of a
// agent.API.GetEntities call.
type AgentGetEntitiesResults struct {
	Entities []AgentGetEntitiesResult `json:"entities"`
}

// AgentGetEntitiesResult holds the results of a
// machineagent.API.GetEntities call for a single entity.
type AgentGetEntitiesResult struct {
	Life  Life   `json:"life"`
	Error *Error `json:"error,omitempty"`
}

// VersionResult holds the version and possibly error for a given
// DesiredVersion() API call.
type VersionResult struct {
	Version *version.Number `json:"version,omitempty"`
	Error   *Error          `json:"error,omitempty"`
}

// VersionResults is a list of versions for the requested entities.
type VersionResults struct {
	Results []VersionResult `json:"results"`
}

// Version holds a specific binary version.
type Version struct {
	Version version.Binary `json:"version"`
}

// EntityVersion specifies the tools version to be set for an entity
// with the given tag.
// version.Binary directly.
type EntityVersion struct {
	Tag string `json:"tag"`
}

// EntitiesVersion specifies what tools are being run for
// multiple entities.
type EntitiesVersion struct {
}

// NotifyWatchResult holds a NotifyWatcher id and an error (if any).
type NotifyWatchResult struct {
	NotifyWatcherId string
	Error           *Error `json:"error,omitempty"`
}

// NotifyWatchResults holds the results for any API call which ends up
// returning a list of NotifyWatchers
type NotifyWatchResults struct {
	Results []NotifyWatchResult `json:"results"`
}

// StringsWatchResult holds a StringsWatcher id, changes and an error
// (if any).
type StringsWatchResult struct {
	StringsWatcherId string   `json:"watcher-id"`
	Changes          []string `json:"changes,omitempty"`
	Error            *Error   `json:"error,omitempty"`
}

// StringsWatchResults holds the results for any API call which ends up
// returning a list of StringsWatchers.
type StringsWatchResults struct {
	Results []StringsWatchResult `json:"results"`
}

// EntitiesWatchResult holds a EntitiesWatcher id, changes and an error
// (if any).
type EntitiesWatchResult struct {
	// Note legacy serialization tag.
	EntitiesWatcherId string   `json:"watcher-id"`
	Changes           []string `json:"changes,omitempty"`
	Error             *Error   `json:"error,omitempty"`
}

// EntitiesWatchResults holds the results for any API call which ends up
// returning a list of EntitiesWatchers.
type EntitiesWatchResults struct {
	Results []EntitiesWatchResult `json:"results"`
}

// UnitSettings specifies the version of some unit's settings in some relation.
type UnitSettings struct {
	Version int64 `json:"version"`
}

// AgentVersionResult is used to return the current version number of the
// agent running the API server.
type AgentVersionResult struct {
	Version version.Number `json:"version"`
}

// LogMessage is a structured logging entry.
type LogMessage struct {
	Entity    string    `json:"tag"`
	Timestamp time.Time `json:"ts"`
	Severity  string    `json:"sev"`
	Module    string    `json:"mod"`
	Location  string    `json:"loc"`
	Message   string    `json:"msg"`
}
