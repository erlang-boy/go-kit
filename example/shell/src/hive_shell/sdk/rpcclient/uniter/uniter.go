// Copyright 2013 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package uniter

import (
	"fmt"

	"github.com/automation/errors"
	"gopkg.in/automation/names.v2"

	"api/base"
	"api/common"
	//_apiwatcher "api/watcher"
	"api/params"
	//"automation/common/network"
	_ "watcher"
)

const uniterFacade = "Uniter"

// State provides access to the Uniter API facade.
type State struct {
	facade base.FacadeCaller
	// unitTag contains the authenticated unit's tag.
	unitTag names.UnitTag
}

// newStateForVersion creates a new client-side Uniter facade for the
// given version.
func newStateForVersion(
	caller base.APICaller,
	authTag names.UnitTag,
	version int,
) *State {
	facadeCaller := base.NewFacadeCallerForVersion(
		caller,
		uniterFacade,
		version,
	)
	state := &State{
		facade:  facadeCaller,
		unitTag: authTag,
	}

	//TODO: not used newWatcher
	//newWatcher := func(result params.NotifyWatchResult) watcher.NotifyWatcher {
	//	return apiwatcher.NewNotifyWatcher(caller, result)
	//}
	return state
}

func newStateForVersionFn(version int) func(base.APICaller, names.UnitTag) *State {
	return func(caller base.APICaller, authTag names.UnitTag) *State {
		return newStateForVersion(caller, authTag, version)
	}
}

// newStateV4 creates a new client-side Uniter facade, version 4.
var newStateV4 = newStateForVersionFn(4)

// NewState creates a new client-side Uniter facade.
// Defined like this to allow patching during tests.
var NewState = newStateV4

// BestAPIVersion returns the API version that we were able to
// determine is supported by both the client and the API Server.
func (st *State) BestAPIVersion() int {
	return st.facade.BestAPIVersion()
}

// Facade returns the current facade.
func (st *State) Facade() base.FacadeCaller {
	return st.facade
}

// life requests the lifecycle of the given entity from the server.
func (st *State) life(tag names.Tag) (params.Life, error) {
	return common.Life(st.facade, tag)
}

// getOneAction retrieves a single Action from the controller.
func (st *State) getOneAction(tag *names.ActionTag) (params.ActionResult, error) {
	nothing := params.ActionResult{}

	args := params.Entities{
		Entities: []params.Entity{
			{Tag: tag.String()},
		},
	}

	var results params.ActionResults
	err := st.facade.FacadeCall("Actions", args, &results)
	if err != nil {
		return nothing, err
	}

	if len(results.Results) > 1 {
		return nothing, fmt.Errorf("expected only 1 action query result, got %d", len(results.Results))
	}

	// handle server errors
	result := results.Results[0]
	if err := result.Error; err != nil {
		return nothing, err
	}

	return result, nil
}

// Unit provides access to methods of a state.Unit through the facade.
func (st *State) Unit(tag names.UnitTag) (*Unit, error) {
	life, err := st.life(tag)
	if err != nil {
		return nil, err
	}
	return &Unit{
		tag:  tag,
		life: life,
		st:   st,
	}, nil
}

// Action returns the Action with the given tag.
func (st *State) Action(tag names.ActionTag) (*Action, error) {
	result, err := st.getOneAction(&tag)
	if err != nil {
		return nil, err
	}
	return &Action{
		name:   result.Action.Name,
		params: result.Action.Parameters,
	}, nil
}

// ActionBegin marks an action as running.
func (st *State) ActionBegin(tag names.ActionTag) error {
	var outcome params.ErrorResults

	args := params.Entities{
		Entities: []params.Entity{
			{Tag: tag.String()},
		},
	}

	err := st.facade.FacadeCall("BeginActions", args, &outcome)
	if err != nil {
		return err
	}
	if len(outcome.Results) != 1 {
		return fmt.Errorf("expected 1 result, got %d", len(outcome.Results))
	}
	result := outcome.Results[0]
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// ActionFinish captures the structured output of an action.
func (st *State) ActionFinish(tag names.ActionTag, status string, results map[string]interface{}, message string) error {
	var outcome params.ErrorResults

	args := params.ActionExecutionResults{
		Results: []params.ActionExecutionResult{
			{
				ActionTag: tag.String(),
				Status:    status,
				Results:   results,
				Message:   message,
			},
		},
	}

	err := st.facade.FacadeCall("FinishActions", args, &outcome)
	if err != nil {
		return err
	}
	if len(outcome.Results) != 1 {
		return fmt.Errorf("expected 1 result, got %d", len(outcome.Results))
	}
	result := outcome.Results[0]
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// ProviderType returns a provider type used by the current automation model.
//
// TODO(dimitern): We might be able to drop this, once we have machine
// addresses implemented fully. See also LP bug 1221798.
func (st *State) ProviderType() (string, error) {
	var result params.StringResult
	err := st.facade.FacadeCall("ProviderType", nil, &result)
	if err != nil {
		return "", err
	}
	if err := result.Error; err != nil {
		return "", err
	}
	return result.Result, nil
}

// ErrIfNotVersionFn returns a function which can be used to check for
// the minimum supported version, and, if appropriate, generate an
// error.
func ErrIfNotVersionFn(minVersion int, bestAPIVersion int) func(string) error {
	return func(fnName string) error {
		if minVersion <= bestAPIVersion {
			return nil
		}
		return errors.NotImplementedf("%s(...) requires v%d+", fnName, minVersion)
	}
}
