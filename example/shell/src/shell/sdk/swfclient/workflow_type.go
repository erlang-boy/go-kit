package swfclient

import (
	"automation/sdk/sdkerr"
)

/**********************
**	register workflow type
***********************/

// Represents a workflow type.
type WorkflowType struct {
	Name    *string `json:"name" min:"1" type:"string" required:"true"`
	Version *string `json:"version" min:"1" type:"string" required:"true"`
}

// SetName sets the Name field's value.
func (s *WorkflowType) SetName(v string) *WorkflowType {
	s.Name = &v
	return s
}

// SetVersion sets the Version field's value.
func (s *WorkflowType) SetVersion(v string) *WorkflowType {
	s.Version = &v
	return s
}
func (s *WorkflowType) Validate() error {
	invalidParams := sdkerr.ErrInvalidParams{Context: "WorkflowType"}
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

type RegisterWorkflowTypeInput struct {
	DefaultChildPolicy                  *string   `json:"defaultChildPolicy" type:"string" enum:"ChildPolicy"`
	DefaultExecutionStartToCloseTimeout *string   `json:"defaultExecutionStartToCloseTimeout" type:"string"`
	DefaultLambdaRole                   *string   `json:"defaultLambdaRole" min:"1" type:"string"`
	DefaultTaskList                     *TaskList `json:"defaultTaskList" type:"structure"`
	DefaultTaskPriority                 *string   `json:"defaultTaskPriority" type:"string"`
	DefaultTaskStartToCloseTimeout      *string   `json:"defaultTaskStartToCloseTimeout" type:"string"`
	Description                         *string   `json:"description" type:"string"`
	Domain                              *string   `json:"domain" min:"1" type:"string" required:"true"`
	Name                                *string   `json:"name" min:"1" type:"string" required:"true"`
	Version                             *string   `json:"version" min:"1" type:"string" required:"true"`
}

func (s *RegisterWorkflowTypeInput) SetDefaultChildPolicy(v string) *RegisterWorkflowTypeInput {
	s.DefaultChildPolicy = &v
	return s
}

// SetDefaultExecutionStartToCloseTimeout sets the DefaultExecutionStartToCloseTimeout field's value.
func (s *RegisterWorkflowTypeInput) SetDefaultExecutionStartToCloseTimeout(v string) *RegisterWorkflowTypeInput {
	s.DefaultExecutionStartToCloseTimeout = &v
	return s
}

// SetDefaultLambdaRole sets the DefaultLambdaRole field's value.
func (s *RegisterWorkflowTypeInput) SetDefaultLambdaRole(v string) *RegisterWorkflowTypeInput {
	s.DefaultLambdaRole = &v
	return s
}

// SetDefaultTaskList sets the DefaultTaskList field's value.
func (s *RegisterWorkflowTypeInput) SetDefaultTaskList(v *TaskList) *RegisterWorkflowTypeInput {
	s.DefaultTaskList = v
	return s
}

// SetDefaultTaskPriority sets the DefaultTaskPriority field's value.
func (s *RegisterWorkflowTypeInput) SetDefaultTaskPriority(v string) *RegisterWorkflowTypeInput {
	s.DefaultTaskPriority = &v
	return s
}

// SetDefaultTaskStartToCloseTimeout sets the DefaultTaskStartToCloseTimeout field's value.
func (s *RegisterWorkflowTypeInput) SetDefaultTaskStartToCloseTimeout(v string) *RegisterWorkflowTypeInput {
	s.DefaultTaskStartToCloseTimeout = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *RegisterWorkflowTypeInput) SetDescription(v string) *RegisterWorkflowTypeInput {
	s.Description = &v
	return s
}

// SetDomain sets the Domain field's value.
func (s *RegisterWorkflowTypeInput) SetDomain(v string) *RegisterWorkflowTypeInput {
	s.Domain = &v
	return s
}

// SetName sets the Name field's value.
func (s *RegisterWorkflowTypeInput) SetName(v string) *RegisterWorkflowTypeInput {
	s.Name = &v
	return s
}

// SetVersion sets the Version field's value.
func (s *RegisterWorkflowTypeInput) SetVersion(v string) *RegisterWorkflowTypeInput {
	s.Version = &v
	return s
}

func (s *RegisterWorkflowTypeInput) Validate() error {
	invalidParams := sdkerr.ErrInvalidParams{Context: "RegisterWorkflowTypeInput"}
	if s.DefaultLambdaRole != nil && len(*s.DefaultLambdaRole) < 1 {
		invalidParams.Add(sdkerr.NewErrParamMinLen("DefaultLambdaRole", 1))
	}
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

type RegisterWorkflowTypeOutput struct {
}

/********************
* list workflow type
********************/
type ListWorkflowTypesInput struct {
	Domain             *string `json:"domain" min:"1" type:"string" required:"true"`
	MaximumPageSize    *int64  `json:"maximumPageSize,omitempty" type:"integer"`
	Name               *string `json:"name,omitempty" min:"1" type:"string"`
	NextPageToken      *string `json:"nextPageToken,omitempty" type:"string"`
	RegistrationStatus *string `json:"registrationStatus,omitempty" type:"string" required:"true" enum:"RegistrationStatus"`
	ReverseOrder       *bool   `json:"reverseOrder,omitempty" type:"boolean"`
}

func (s *ListWorkflowTypesInput) SetDomain(v string) *ListWorkflowTypesInput {
	s.Domain = &v
	return s
}

// SetMaximumPageSize sets the MaximumPageSize field's value.
func (s *ListWorkflowTypesInput) SetMaximumPageSize(v int64) *ListWorkflowTypesInput {
	s.MaximumPageSize = &v
	return s
}

// SetName sets the Name field's value.
func (s *ListWorkflowTypesInput) SetName(v string) *ListWorkflowTypesInput {
	s.Name = &v
	return s
}

// SetNextPageToken sets the NextPageToken field's value.
func (s *ListWorkflowTypesInput) SetNextPageToken(v string) *ListWorkflowTypesInput {
	s.NextPageToken = &v
	return s
}

// SetRegistrationStatus sets the RegistrationStatus field's value.
func (s *ListWorkflowTypesInput) SetRegistrationStatus(v string) *ListWorkflowTypesInput {
	s.RegistrationStatus = &v
	return s
}

// SetReverseOrder sets the ReverseOrder field's value.
func (s *ListWorkflowTypesInput) SetReverseOrder(v bool) *ListWorkflowTypesInput {
	s.ReverseOrder = &v
	return s
}

// Contains a paginated list of information structures about workflow types.

func (s *ListWorkflowTypesInput) Validate() error {
	invalidParams := sdkerr.ErrInvalidParams{Context: "ListWorkflowTypesInput"}
	if s.Domain == nil {
		invalidParams.Add(sdkerr.NewErrParamRequired("Domain"))
	}
	if s.Domain != nil && len(*s.Domain) < 1 {
		invalidParams.Add(sdkerr.NewErrParamMinLen("Domain", 1))
	}
	//	if s.Name != nil && len(*s.Name) < 1 {
	//		invalidParams.Add(sdkerr.NewErrParamMinLen("Name", 1))
	//	}
	if s.RegistrationStatus == nil {
		invalidParams.Add(sdkerr.NewErrParamRequired("RegistrationStatus"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type WorkflowTypeInfo struct {
	CreationDate    *int64        `json:"creationDate" type:"timestamp" timestampFormat:"unix" required:"true"`
	DeprecationDate *int64        `json:"deprecationDate" type:"timestamp" timestampFormat:"unix"`
	Description     *string       `json:"description" type:"string"`
	Status          *string       `json:"status" type:"string" required:"true" enum:"RegistrationStatus"`
	WorkflowType    *WorkflowType `json:"workflowType" type:"structure" required:"true"`
}

type ListWorkflowTypesOutput struct {
	NextPageToken *string             `json:"nextPageToken" type:"string"`
	TypeInfos     []*WorkflowTypeInfo `json:"typeInfos" type:"list" required:"true"`
}

/*********
*  describe workflow type
**********/

type DescribeWorkflowTypeInput struct {
	Domain       *string       `json:"domain" type:"string" required:"true"`
	WorkflowType *WorkflowType `json:"workflowType" type:"structure" required:"true"`
}

func (d *DescribeWorkflowTypeInput) SetDomain(domain string) *DescribeWorkflowTypeInput {
	d.Domain = &domain
	return d
}

func (d *DescribeWorkflowTypeInput) SetWorkflowType(w *WorkflowType) *DescribeWorkflowTypeInput {
	d.WorkflowType = w
	return d
}

func (d *DescribeWorkflowTypeInput) Validate() error {
	invalidParams := sdkerr.ErrInvalidParams{Context: "DescribeWorkflowTypeInput"}
	if d.Domain == nil {
		invalidParams.Add(sdkerr.NewErrParamRequired("Domain"))
	}

	if d.Domain != nil && len(*d.Domain) < 1 {
		invalidParams.Add(sdkerr.NewErrParamMinLen("Domain", 1))
	}

	if d.WorkflowType == nil {
		invalidParams.Add(sdkerr.NewErrParamRequired("WorkflowType"))
	}

	if d.WorkflowType != nil {
		if err := d.WorkflowType.Validate(); err != nil {
			invalidParams.AddNested("WorkflowType", err.(sdkerr.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeWorkflowTypeOutput struct {
	Configuration *WorkflowConfiguration `json:"configuration,omitempty" type:"structure" required:"true"`
	TypeInfo      *WorkflowTypeInfo      `json:"typeInfo,omitempty" type:"structure" required:"true"`
}

type WorkflowConfiguration struct {
	DefaultChildPolicy                  string    `json:"defaultChildPolicy,omitempty" type:"string"`
	DefaultExecutionStartToCloseTimeout string    `json:"defaultExecutionStartToCloseTimeout,omitempty" type:"string"`
	DefaultLambdaRole                   string    `json:"defaultLambdaRole,omitempty" type:"string"`
	DefaultTaskList                     *TaskList `json:"defaultTaskList,omitempty" type:"structure"`
	DefaultTaskPriority                 string    `json:"defaultTaskPriority,omitempty" type:"string"`
	DefaultTaskStartToCloseTimeout      string    `json:"defaultTaskStartToCloseTimeout,omitempty" type:"string"`
}
