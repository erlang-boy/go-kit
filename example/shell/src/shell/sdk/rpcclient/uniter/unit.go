// Copyright 2013 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package uniter

import (
	"fmt"

	"github.com/automation/errors"
	"gopkg.in/automation/names.v2"

	"api/common"
	"api/params"
	apiwatcher "api/watcher"
	"status"
	"watcher"
)

type Unit struct {
	st   *State
	tag  names.UnitTag
	life params.Life
}

// Tag returns the unit's tag.
func (u *Unit) Tag() names.UnitTag {
	return u.tag
}

// Name returns the name of the unit.
func (u *Unit) Name() string {
	return u.tag.Id()
}

// String returns the unit as a string.
func (u *Unit) String() string {
	return u.Name()
}

// Life returns the unit's lifecycle value.
func (u *Unit) Life() params.Life {
	return u.life
}

// Refresh updates the cached local copy of the unit's data.
func (u *Unit) Refresh() error {
	life, err := u.st.life(u.tag)
	if err != nil {
		return err
	}
	u.life = life
	return nil
}

// SetUnitStatus sets the status of the unit.
func (u *Unit) SetUnitStatus(unitStatus status.Status, info string, data map[string]interface{}) error {
	if u.st.facade.BestAPIVersion() < 2 {
		return errors.NotImplementedf("SetUnitStatus")
	}
	var result params.ErrorResults
	args := params.SetStatus{
		Entities: []params.EntityStatusArgs{
			{Tag: u.tag.String(), Status: unitStatus.String(), Info: info, Data: data},
		},
	}
	err := u.st.facade.FacadeCall("SetUnitStatus", args, &result)
	if err != nil {
		return errors.Trace(err)
	}
	return result.OneError()
}

// UnitStatus gets the status details of the unit.
func (u *Unit) UnitStatus() (params.StatusResult, error) {
	var results params.StatusResults
	args := params.Entities{
		Entities: []params.Entity{
			{Tag: u.tag.String()},
		},
	}
	err := u.st.facade.FacadeCall("UnitStatus", args, &results)
	if err != nil {
		return params.StatusResult{}, errors.Trace(err)
	}
	if len(results.Results) != 1 {
		panic(errors.Errorf("expected 1 result, got %d", len(results.Results)))
	}
	result := results.Results[0]
	if result.Error != nil {
		return params.StatusResult{}, result.Error
	}
	return result, nil
}

// SetAgentStatus sets the status of the unit agent.
func (u *Unit) SetAgentStatus(agentStatus status.Status, info string, data map[string]interface{}) error {
	var result params.ErrorResults
	args := params.SetStatus{
		Entities: []params.EntityStatusArgs{
			{Tag: u.tag.String(), Status: agentStatus.String(), Info: info, Data: data},
		},
	}
	setStatusFacadeCall := "SetAgentStatus"
	if u.st.facade.BestAPIVersion() < 2 {
		setStatusFacadeCall = "SetStatus"
	}
	err := u.st.facade.FacadeCall(setStatusFacadeCall, args, &result)
	if err != nil {
		return err
	}
	return result.OneError()
}

// EnsureDead sets the unit lifecycle to Dead if it is Alive or
// Dying. It does nothing otherwise.
func (u *Unit) EnsureDead() error {
	var result params.ErrorResults
	args := params.Entities{
		Entities: []params.Entity{{Tag: u.tag.String()}},
	}
	err := u.st.facade.FacadeCall("EnsureDead", args, &result)
	if err != nil {
		return err
	}
	return result.OneError()
}

//patched by vega
// EnsureHalted sets the unit lifecycle to halted if it is Alive or
// Dying. It does nothing otherwise.
func (u *Unit) EnsureHalted() error {
	var result params.ErrorResults
	args := params.Entities{
		Entities: []params.Entity{{Tag: u.tag.String()}},
	}
	err := u.st.facade.FacadeCall("EnsureHalted", args, &result)
	if err != nil {
		return err
	}
	return result.OneError()
}

// Watch returns a watcher for observing changes to the unit.
func (u *Unit) Watch() (watcher.NotifyWatcher, error) {
	return common.Watch(u.st.facade, u.tag)
}

// Destroy, when called on a Alive unit, advances its lifecycle as far as
// possible; it otherwise has no effect. In most situations, the unit's
// life is just set to Dying; but if a principal unit that is not assigned
// to a provisioned machine is Destroyed, it will be removed from state
// directly.
func (u *Unit) Destroy() error {
	var result params.ErrorResults
	args := params.Entities{
		Entities: []params.Entity{{Tag: u.tag.String()}},
	}
	err := u.st.facade.FacadeCall("Destroy", args, &result)
	if err != nil {
		return err
	}
	return result.OneError()
}

// PublicAddress returns the public address of the unit and whether it
// is valid.
//
// NOTE: This differs from state.Unit.PublicAddres() by returning
// an error instead of a bool, because it needs to make an API call.
//
// TODO(dimitern): We might be able to drop this, once we have machine
// addresses implemented fully. See also LP bug 1221798.
func (u *Unit) PublicAddress() (string, error) {
	var results params.StringResults
	args := params.Entities{
		Entities: []params.Entity{{Tag: u.tag.String()}},
	}
	err := u.st.facade.FacadeCall("PublicAddress", args, &results)
	if err != nil {
		return "", err
	}
	if len(results.Results) != 1 {
		return "", fmt.Errorf("expected 1 result, got %d", len(results.Results))
	}
	result := results.Results[0]
	if result.Error != nil {
		return "", result.Error
	}
	return result.Result, nil
}

// PrivateAddress returns the private address of the unit and whether
// it is valid.
//
// NOTE: This differs from state.Unit.PrivateAddress() by returning
// an error instead of a bool, because it needs to make an API call.
//
// TODO(dimitern): We might be able to drop this, once we have machine
// addresses implemented fully. See also LP bug 1221798.
func (u *Unit) PrivateAddress() (string, error) {
	var results params.StringResults
	args := params.Entities{
		Entities: []params.Entity{{Tag: u.tag.String()}},
	}
	err := u.st.facade.FacadeCall("PrivateAddress", args, &results)
	if err != nil {
		return "", err
	}
	if len(results.Results) != 1 {
		return "", fmt.Errorf("expected 1 result, got %d", len(results.Results))
	}
	result := results.Results[0]
	if result.Error != nil {
		return "", result.Error
	}
	return result.Result, nil
}

// WatchConfigSettings returns a watcher for observing changes to the
// unit's service configuration settings. The unit must have a charm URL
// set before this method is called, and the returned watcher will be
// valid only while the unit's charm URL is not changed.
func (u *Unit) WatchConfigSettings() (watcher.NotifyWatcher, error) {
	var results params.NotifyWatchResults
	args := params.Entities{
		Entities: []params.Entity{{Tag: u.tag.String()}},
	}
	err := u.st.facade.FacadeCall("WatchConfigSettings", args, &results)
	if err != nil {
		return nil, err
	}
	if len(results.Results) != 1 {
		return nil, fmt.Errorf("expected 1 result, got %d", len(results.Results))
	}
	result := results.Results[0]
	if result.Error != nil {
		return nil, result.Error
	}
	w := apiwatcher.NewNotifyWatcher(u.st.facade.RawAPICaller(), result)
	return w, nil
}

// WatchAddresses returns a watcher for observing changes to the
// unit's addresses. The unit must be assigned to a machine before
// this method is called, and the returned watcher will be valid only
// while the unit's assigned machine is not changed.
func (u *Unit) WatchAddresses() (watcher.NotifyWatcher, error) {
	var results params.NotifyWatchResults
	args := params.Entities{
		Entities: []params.Entity{{Tag: u.tag.String()}},
	}
	err := u.st.facade.FacadeCall("WatchUnitAddresses", args, &results)
	if err != nil {
		return nil, err
	}
	if len(results.Results) != 1 {
		return nil, fmt.Errorf("expected 1 result, got %d", len(results.Results))
	}
	result := results.Results[0]
	if result.Error != nil {
		return nil, result.Error
	}
	w := apiwatcher.NewNotifyWatcher(u.st.facade.RawAPICaller(), result)
	return w, nil
}
