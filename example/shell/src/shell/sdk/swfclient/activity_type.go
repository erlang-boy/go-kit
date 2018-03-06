package swfclient

import (
	"automation/sdk/sdkerr"
	"time"
)

/**************************
*  register activity type
***************************/

// Represents an activity type.
type ActivityType struct {
	Name    *string `json:"name" min:"1" type:"string" required:"true"`
	Version *string `json:"version" min:"1" type:"string" required:"true"`
}

// SetName sets the Name field's value.
func (s *ActivityType) SetName(v string) *ActivityType {
	s.Name = &v
	return s
}

// SetVersion sets the Version field's value.
func (s *ActivityType) SetVersion(v string) *ActivityType {
	s.Version = &v
	return s
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ActivityType) Validate() error {
	invalidParams := sdkerr.ErrInvalidParams{Context: "ActivityType"}
	if s.Name == nil {
		invalidParams.Add(sdkerr.NewErrParamRequired("Name"))
	}
	if s.Name != nil && len(*s.Name) < 1 {
		invalidParams.Add(sdkerr.NewErrParamMinLen("Name", 1))
	}
	if s.Version == nil {
		invalidParams.Add(sdkerr.NewErrParamRequired("Version"))
	}
	if s.Version != nil && len(*s.Version) < 1 {
		invalidParams.Add(sdkerr.NewErrParamMinLen("Version", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type RegisterActivityTypeInput struct {
	DefaultTaskHeartbeatTimeout       *string   `json:"defaultTaskHeartbeatTimeout,omitempty" type:"string"`
	DefaultTaskList                   *TaskList `json:"defaultTaskList,omitempty" type:"structure"`
	DefaultTaskPriority               *string   `json:"defaultTaskPriority,omitempty" type:"string"`
	DefaultTaskScheduleToCloseTimeout *string   `json:"defaultTaskScheduleToCloseTimeout,omitempty" type:"string"`
	DefaultTaskScheduleToStartTimeout *string   `json:"defaultTaskScheduleToStartTimeout,omitempty" type:"string"`
	DefaultTaskStartToCloseTimeout    *string   `json:"defaultTaskStartToCloseTimeout,omitempty" type:"string"`
	Description                       *string   `json:"description,omitempty" type:"string"`
	Domain                            *string   `json:"domain" min:"1" type:"string" required:"true"`
	Name                              *string   `json:"name" min:"1" type:"string" required:"true"`
	Version                           *string   `json:"version" min:"1" type:"string" required:"true"`
}

func (s *RegisterActivityTypeInput) SetDefaultTaskHeartbeatTimeout(v string) *RegisterActivityTypeInput {
	s.DefaultTaskHeartbeatTimeout = &v
	return s
}

// SetDefaultTaskList sets the DefaultTaskList field's value.
func (s *RegisterActivityTypeInput) SetDefaultTaskList(v *TaskList) *RegisterActivityTypeInput {
	s.DefaultTaskList = v
	return s
}

// SetDefaultTaskPriority sets the DefaultTaskPriority field's value.
func (s *RegisterActivityTypeInput) SetDefaultTaskPriority(v string) *RegisterActivityTypeInput {
	s.DefaultTaskPriority = &v
	return s
}

// SetDefaultTaskScheduleToCloseTimeout sets the DefaultTaskScheduleToCloseTimeout field's value.
func (s *RegisterActivityTypeInput) SetDefaultTaskScheduleToCloseTimeout(v string) *RegisterActivityTypeInput {
	s.DefaultTaskScheduleToCloseTimeout = &v
	return s
}

// SetDefaultTaskScheduleToStartTimeout sets the DefaultTaskScheduleToStartTimeout field's value.
func (s *RegisterActivityTypeInput) SetDefaultTaskScheduleToStartTimeout(v string) *RegisterActivityTypeInput {
	s.DefaultTaskScheduleToStartTimeout = &v
	return s
}

// SetDefaultTaskStartToCloseTimeout sets the DefaultTaskStartToCloseTimeout field's value.
func (s *RegisterActivityTypeInput) SetDefaultTaskStartToCloseTimeout(v string) *RegisterActivityTypeInput {
	s.DefaultTaskStartToCloseTimeout = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *RegisterActivityTypeInput) SetDescription(v string) *RegisterActivityTypeInput {
	s.Description = &v
	return s
}

// SetDomain sets the Domain field's value.
func (s *RegisterActivityTypeInput) SetDomain(v string) *RegisterActivityTypeInput {
	s.Domain = &v
	return s
}

// SetName sets the Name field's value.
func (s *RegisterActivityTypeInput) SetName(v string) *RegisterActivityTypeInput {
	s.Name = &v
	return s
}

// SetVersion sets the Version field's value.
func (s *RegisterActivityTypeInput) SetVersion(v string) *RegisterActivityTypeInput {
	s.Version = &v
	return s
}

func (s *RegisterActivityTypeInput) Validate() error {
	invalidParams := sdkerr.ErrInvalidParams{Context: "RegisterActivityTypeInput"}
	if s.Domain == nil {
		invalidParams.Add(sdkerr.NewErrParamRequired("Domain"))
	}
	if s.Domain != nil && len(*s.Domain) < 1 {
		invalidParams.Add(sdkerr.NewErrParamMinLen("Domain", 1))
	}
	if s.Name == nil {
		invalidParams.Add(sdkerr.NewErrParamRequired("Name"))
	}
	if s.Name != nil && len(*s.Name) < 1 {
		invalidParams.Add(sdkerr.NewErrParamMinLen("Name", 1))
	}
	if s.Version == nil {
		invalidParams.Add(sdkerr.NewErrParamRequired("Version"))
	}
	if s.Version != nil && len(*s.Version) < 1 {
		invalidParams.Add(sdkerr.NewErrParamMinLen("Version", 1))
	}
	if s.DefaultTaskList != nil {
		if err := s.DefaultTaskList.Validate(); err != nil {
			invalidParams.AddNested("DefaultTaskList", err.(sdkerr.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type RegisterActivityTypeOutput struct {
}

/******************
* list activity type
********************/
type ListActivityTypesInput struct {
	Domain             *string `json:"domain" min:"1" type:"string" required:"true"`
	MaximumPageSize    *int64  `json:"maximumPageSize,omitempty" type:"integer"`
	Name               *string `json:"name,omitempty" min:"1" type:"string"`
	NextPageToken      *string `json:"nextPageToken,omitempty" type:"string"`
	RegistrationStatus *string `json:"registrationStatus" type:"string" required:"true" enum:"RegistrationStatus"`
	ReverseOrder       *bool   `json:"reverseOrder,omitempty" type:"boolean"`
}

func (s *ListActivityTypesInput) SetDomain(v string) *ListActivityTypesInput {
	s.Domain = &v
	return s
}
func (s *ListActivityTypesInput) SetMaximumPageSize(v int64) *ListActivityTypesInput {
	s.MaximumPageSize = &v
	return s
}
func (s *ListActivityTypesInput) SetName(v string) *ListActivityTypesInput {
	s.Name = &v
	return s
}

// SetNextPageToken sets the NextPageToken field's value.
func (s *ListActivityTypesInput) SetNextPageToken(v string) *ListActivityTypesInput {
	s.NextPageToken = &v
	return s
}

// SetRegistrationStatus sets the RegistrationStatus field's value.
func (s *ListActivityTypesInput) SetRegistrationStatus(v string) *ListActivityTypesInput {
	s.RegistrationStatus = &v
	return s
}

// SetReverseOrder sets the ReverseOrder field's value.
func (s *ListActivityTypesInput) SetReverseOrder(v bool) *ListActivityTypesInput {
	s.ReverseOrder = &v
	return s
}

type ActivityTypeInfo struct {
	ActivityType    *ActivityType `json:"activityType" type:"structure" required:"true"`
	CreationDate    *time.Time    `json:"creationDate" type:"timestamp" timestampFormat:"unix" required:"true"`
	DeprecationDate *time.Time    `json:"deprecationDate" type:"timestamp" timestampFormat:"unix"`
	Description     *string       `json:"description" type:"string"`
	Status          *string       `json:"status" type:"string" required:"true" enum:"RegistrationStatus"`
}

func (s *ListActivityTypesInput) Validate() error {
	invalidParams := sdkerr.ErrInvalidParams{Context: "ListActivityTypesInput"}
	if s.Domain == nil {
		invalidParams.Add(sdkerr.NewErrParamRequired("Domain"))
	}
	if s.Domain != nil && len(*s.Domain) < 1 {
		invalidParams.Add(sdkerr.NewErrParamMinLen("Domain", 1))
	}
	if s.Name != nil && len(*s.Name) < 1 {
		invalidParams.Add(sdkerr.NewErrParamMinLen("Name", 1))
	}
	if s.RegistrationStatus == nil {
		invalidParams.Add(sdkerr.NewErrParamRequired("RegistrationStatus"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Contains a paginated list of activity type information structures.
type ListActivityTypesOutput struct {
	NextPageToken *string             `json:"nextPageToken" type:"string"`
	TypeInfos     []*ActivityTypeInfo `json:"typeInfos" type:"list" required:"true"`
}

/************************
* describe activity type
************************/

type DescribeActivityTypeInput struct {
	ActivityType *ActivityType `json:"activityType" type:"structure" required:"true"`
	Domain       *string       `json:"domain" min:"1" type:"string" required:"true"`
}

// SetActivityType sets the ActivityType field's value.
func (s *DescribeActivityTypeInput) SetActivityType(v *ActivityType) *DescribeActivityTypeInput {
	s.ActivityType = v
	return s
}

// SetDomain sets the Domain field's value.
func (s *DescribeActivityTypeInput) SetDomain(v string) *DescribeActivityTypeInput {
	s.Domain = &v
	return s
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeActivityTypeInput) Validate() error {
	invalidParams := sdkerr.ErrInvalidParams{Context: "DescribeActivityTypeInput"}
	if s.ActivityType == nil {
		invalidParams.Add(sdkerr.NewErrParamRequired("ActivityType"))
	}
	if s.Domain == nil {
		invalidParams.Add(sdkerr.NewErrParamRequired("Domain"))
	}
	if s.Domain != nil && len(*s.Domain) < 1 {
		invalidParams.Add(sdkerr.NewErrParamMinLen("Domain", 1))
	}

	if s.ActivityType != nil {
		if err := s.ActivityType.Validate(); err != nil {
			invalidParams.AddNested("ActivityType", err.(sdkerr.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Configuration settings registered with the activity type.
type ActivityTypeConfiguration struct {
	DefaultTaskHeartbeatTimeout *string `json:"defaultTaskHeartbeatTimeout,omitempty" type:"string"`

	DefaultTaskList *TaskList `json:"defaultTaskList,omitempty" type:"structure"`

	DefaultTaskPriority *string `json:"defaultTaskPriority,omitempty" type:"string"`

	DefaultTaskScheduleToCloseTimeout *string `json:"defaultTaskScheduleToCloseTimeout,omitempty" type:"string"`

	DefaultTaskScheduleToStartTimeout *string `json:"defaultTaskScheduleToStartTimeout,omitempty" type:"string"`

	DefaultTaskStartToCloseTimeout *string `json:"defaultTaskStartToCloseTimeout,omitempty" type:"string"`
}

// Detailed information about an activity type.
type DescribeActivityTypeOutput struct {
	Configuration *ActivityTypeConfiguration `json:"configuration" type:"structure" required:"true"`

	TypeInfo *ActivityTypeInfo `json:"typeInfo" type:"structure" required:"true"`
}
