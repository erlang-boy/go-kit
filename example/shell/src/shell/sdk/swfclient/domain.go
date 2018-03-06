package swfclient

import (
	"automation/sdk/sdkerr"
)

/************************************
*   register domain
**/

type RegisterDomainInput struct {
	Description *string `json:"description,omitempty" type:"string"`
	Name        *string `json:"name,omitempty" min:"1" type:"string" required:"true"`

	// The duration (in days) that records and histories of workflow executions
	WorkflowExecutionRetentionPeriodInDays *string `json:"workflowExecutionRetentionPeriodInDays,omitempty" min:"1" type:"string" required:"true"`
}

func (s *RegisterDomainInput) SetDescription(v string) *RegisterDomainInput {
	s.Description = &v
	return s
}

func (s *RegisterDomainInput) SetName(v string) *RegisterDomainInput {
	s.Name = &v
	return s
}

func (s *RegisterDomainInput) SetWorkflowExecutionRetentionPeriodInDays(v string) *RegisterDomainInput {
	s.WorkflowExecutionRetentionPeriodInDays = &v
	return s
}

func (s *RegisterDomainInput) Validate() error {
	invalidParams := sdkerr.ErrInvalidParams{Context: "RegisterDomainInput"}
	if s.Name == nil {
		invalidParams.Add(sdkerr.NewErrParamRequired("Name"))
	}
	if s.Name != nil && len(*s.Name) < 1 {
		invalidParams.Add(sdkerr.NewErrParamMinLen("Name", 1))
	}
	if s.WorkflowExecutionRetentionPeriodInDays == nil {
		invalidParams.Add(sdkerr.NewErrParamRequired("WorkflowExecutionRetentionPeriodInDays"))
	}
	if s.WorkflowExecutionRetentionPeriodInDays != nil && len(*s.WorkflowExecutionRetentionPeriodInDays) < 1 {
		invalidParams.Add(sdkerr.NewErrParamMinLen("WorkflowExecutionRetentionPeriodInDays", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type RegisterDomainOutput struct {
}

/************************************
*   list domains
 */
type ListDomainsInput struct {
	MaximumPageSize    *int64  `json:"maximumPageSize,omitempty" type:"integer"`
	NextPageToken      *string `json:"nextPageToken,omitempty" type:"string"`
	RegistrationStatus *string `json:"registrationStatus,omitempty" type:"string" required:"true" enum:"RegistrationStatus"`
	ReverseOrder       *bool   `json:"reverseOrder,omitempty" type:"boolean"`
}

// SetMaximumPageSize sets the MaximumPageSize field's value.
func (s *ListDomainsInput) SetMaximumPageSize(v int64) *ListDomainsInput {
	s.MaximumPageSize = &v
	return s
}

// SetNextPageToken sets the NextPageToken field's value.
func (s *ListDomainsInput) SetNextPageToken(v string) *ListDomainsInput {
	s.NextPageToken = &v
	return s
}

// SetRegistrationStatus sets the RegistrationStatus field's value.
func (s *ListDomainsInput) SetRegistrationStatus(v string) *ListDomainsInput {
	s.RegistrationStatus = &v
	return s
}

// SetReverseOrder sets the ReverseOrder field's value.
func (s *ListDomainsInput) SetReverseOrder(v bool) *ListDomainsInput {
	s.ReverseOrder = &v
	return s
}
func (s *ListDomainsInput) Validate() error {
	invalidParams := sdkerr.ErrInvalidParams{Context: "ListDomainsInput"}
	if s.RegistrationStatus == nil {
		invalidParams.Add(sdkerr.NewErrParamRequired("RegistrationStatus"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ListDomainsOutput struct {
	DomainInfos   []*DomainInfo `json:"domainInfos,omitempty" type:"list" required:"true"`
	NextPageToken *string       `json:"nextPageToken,omitempty" type:"string"`
}

type DomainInfo struct {
	Description *string `json:"description,omitempty" type:"string"`
	Name        *string `json:"name,omitempty" min:"1" type:"string" required:"true"`
	Status      *string `json:"status,omitempty" type:"string" required:"true" enum:"RegistrationStatus"`
}

/************************************
* describe domain
 */

type DomainConfiguration struct {
	WorkflowExecutionRetentionPeriodInDays *string `json:"workflowExecutionRetentionPeriodInDays,omitempty" min:"1" type:"string" required:"true"`
}

type DescribeDomainInput struct {
	// The name of the domain to describe.
	Name *string `json:"name,omitempty" min:"1" type:"string" required:"true"`
}

// SetName sets the Name field's value.
func (s *DescribeDomainInput) SetName(v string) *DescribeDomainInput {
	s.Name = &v
	return s
}

func (s *DescribeDomainInput) Validate() error {
	invalidParams := sdkerr.ErrInvalidParams{Context: "DescribeDomainInput"}
	if s.Name == nil {
		invalidParams.Add(sdkerr.NewErrParamRequired("Name"))
	}
	if s.Name != nil && len(*s.Name) < 1 {
		invalidParams.Add(sdkerr.NewErrParamMinLen("Name", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeDomainOutput struct {
	Configuration *DomainConfiguration `json:"configuration,omitempty" type:"structure" required:"true"`
	DomainInfo    *DomainInfo          `json:"domainInfo,omitempty" type:"structure" required:"true"`
}
