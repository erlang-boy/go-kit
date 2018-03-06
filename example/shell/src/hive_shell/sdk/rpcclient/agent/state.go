// Copyright 2012, 2013 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package agent

import (
	"fmt"

	_ "github.com/automation/errors"
	"gopkg.in/automation/names.v2"

	"api/base"
	_ "api/common"
	//_apiwatcher "api/watcher"
	"api/params"
	_ "watcher"
)

// State provides access to an agent's view of the state.
type State struct {
	facade base.FacadeCaller
}

// NewState returns a version of the state that provides functionality
// required by agent code.
func NewState(caller base.APICaller) *State {
	facadeCaller := base.NewFacadeCaller(caller, "Agent")
	return &State{
		facade: facadeCaller,
	}
}

func (st *State) getEntity(tag names.Tag) (*params.AgentGetEntitiesResult, error) {
	var results params.AgentGetEntitiesResults
	args := params.Entities{
		Entities: []params.Entity{{Tag: tag.String()}},
	}
	err := st.facade.FacadeCall("GetEntities", args, &results)
	if err != nil {
		return nil, err
	}
	if len(results.Entities) != 1 {
		return nil, fmt.Errorf("expected 1 result, got %d", len(results.Entities))
	}
	if err := results.Entities[0].Error; err != nil {
		return nil, err
	}
	return &results.Entities[0], nil
}

type Entity struct {
	st  *State
	tag names.Tag
	doc params.AgentGetEntitiesResult
}

func (st *State) Entity(tag names.Tag) (*Entity, error) {
	doc, err := st.getEntity(tag)
	if err != nil {
		return nil, err
	}
	return &Entity{
		st:  st,
		tag: tag,
		doc: *doc,
	}, nil
}

// Tag returns the entity's tag.
func (m *Entity) Tag() string {
	return m.tag.String()
}

// Life returns the current life cycle state of the entity.
func (m *Entity) Life() params.Life {
	return m.doc.Life
}

// ContainerType returns the type of container hosting this entity.
// If the entity is not a machine, it returns an empty string.
//func (m *Entity) ContainerType() instance.ContainerType {
//	return m.doc.ContainerType
//}

// ClearReboot clears the reboot flag of the machine.
func (m *Entity) ClearReboot() error {
	var result params.ErrorResults
	args := params.SetStatus{
		Entities: []params.EntityStatusArgs{
			{Tag: m.tag.String()},
		},
	}
	err := m.st.facade.FacadeCall("ClearReboot", args, &result)
	if err != nil {
		return err
	}
	return result.OneError()
}
