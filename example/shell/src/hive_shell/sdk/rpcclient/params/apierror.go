// Copyright 2013 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package params

import (
	"fmt"

	"github.com/automation/errors"
)

// UpgradeInProgressError signifies an upgrade is in progress.

// Error is the type of error returned by any call to the state API.
type Error struct {
	Message string     `json:"message"`
	Code    string     `json:"code"`
	Info    *ErrorInfo `json:"info,omitempty"`
}

// ErrorInfo holds additional information provided by an error.
// Note that although these fields are compatible with the
// same fields in httpbakery.ErrorInfo, the Automation API server does
// not implement endpoints directly compatible with that protocol
// because the error response format varies according to
// the endpoint.
type ErrorInfo struct {
}

func (e Error) Error() string {
	return e.Message
}

func (e Error) ErrorCode() string {
	return e.Code
}

// GoString implements fmt.GoStringer.  It means that a *Error shows its
// contents correctly when printed with %#v.
func (e Error) GoString() string {
	return fmt.Sprintf("&params.Error{Message: %q, Code: %q}", e.Message, e.Code)
}

// The Code constants hold error codes for some kinds of error.
const (
	CodeNotFound            = "not found"
	CodeExcessiveContention = "excessive contention"
	CodeNotAssigned         = "not assigned"
	CodeStopped             = "stopped"
	CodeDead                = "dead"
	CodeNoAddressSet        = "no address set"
	CodeTryAgain            = "try again"
	CodeNotImplemented      = "not implemented" // asserted to match rpc.codeNotImplemented in rpc/rpc_test.go
	CodeAlreadyExists       = "already exists"
	CodeActionNotAvailable  = "action no longer available"
	CodeOperationBlocked    = "operation is blocked"
	CodeNotSupported        = "not supported"
	CodeBadRequest          = "bad request"
	CodeMethodNotAllowed    = "method not allowed"
	CodeForbidden           = "forbidden"
	CodeRedirect            = "redirection required"
	CodeRetry               = "retry"
)

// ErrCode returns the error code associated with
// the given error, or the empty string if there
// is none.
func ErrCode(err error) string {
	type ErrorCoder interface {
		ErrorCode() string
	}
	switch err := errors.Cause(err).(type) {
	case ErrorCoder:
		return err.ErrorCode()
	default:
		return ""
	}
}

func IsCodeActionNotAvailable(err error) bool {
	return ErrCode(err) == CodeActionNotAvailable
}

func IsCodeNotFound(err error) bool {
	return ErrCode(err) == CodeNotFound
}

func IsCodeExcessiveContention(err error) bool {
	return ErrCode(err) == CodeExcessiveContention
}

func IsCodeNotAssigned(err error) bool {
	return ErrCode(err) == CodeNotAssigned
}

func IsCodeStopped(err error) bool {
	return ErrCode(err) == CodeStopped
}

func IsCodeDead(err error) bool {
	return ErrCode(err) == CodeDead
}

func IsCodeNoAddressSet(err error) bool {
	return ErrCode(err) == CodeNoAddressSet
}

func IsCodeTryAgain(err error) bool {
	return ErrCode(err) == CodeTryAgain
}

func IsCodeNotImplemented(err error) bool {
	return ErrCode(err) == CodeNotImplemented
}

func IsCodeAlreadyExists(err error) bool {
	return ErrCode(err) == CodeAlreadyExists
}

func IsCodeOperationBlocked(err error) bool {
	return ErrCode(err) == CodeOperationBlocked
}

func IsCodeNotSupported(err error) bool {
	return ErrCode(err) == CodeNotSupported
}

func IsBadRequest(err error) bool {
	return ErrCode(err) == CodeBadRequest
}

func IsMethodNotAllowed(err error) bool {
	return ErrCode(err) == CodeMethodNotAllowed
}

func IsRedirect(err error) bool {
	return ErrCode(err) == CodeRedirect
}
