package errors

import (
	"fmt"
	"net/http"
)

const (
	// BadRequest represents errors caused by wrong query
	// parameters or invalid payload.
	BadRequest Kind = http.StatusBadRequest
	// Unauthorized represents errors connected with auth
	// problems. It is used when handling secure resources calls.
	Unauthorized Kind = http.StatusUnauthorized
	// NotFound represents errors caused by handling calls
	// to nonexistent resources.
	NotFound Kind = http.StatusNotFound
	// Internal represents errors caused by
	// infrastructure malfunction.
	Internal Kind = http.StatusInternalServerError
)

// Kind represents different kinds of possible errors.
type Kind int

// New returns a new Error with the current Kind.
func (k Kind) New(ctx string) *Error {
	return newError(k, ctx)
}

// Newf returns a new Error with the current Kind.
func (k Kind) Newf(ctx string, args ...interface{}) *Error {
	return newfError(k, ctx, args...)
}

// Wrap wraps the cause into an Error with the current Kind.
func (k Kind) Wrap(cause error, ctx string) *Error {
	return wrapError(k, cause, ctx)
}

// Wrapf wraps the cause into an Error with formatting with the current Kind.
func (k Kind) Wrapf(cause error, ctx string, args ...interface{}) *Error {
	return wrapfError(k, cause, ctx, args...)
}

// Error is a wrapper around an error that also contains its context
// and kind.
type Error struct {
	kind  Kind
	ctx   string
	cause error
}

// Error implements the Error interface.
func (e *Error) Error() string {
	return fmt.Sprintf("(%d) %s: %v", e.kind, e.ctx, e.cause)
}

// Kind returns the error kind.
func (e *Error) Kind() Kind {
	return e.kind
}

// Context returns the error context.
func (e *Error) Context() string {
	return e.ctx
}

// Cause returns the error cause.
func (e *Error) Cause() error {
	return e.cause
}

// newError compiles and returns a new instance of the Error.
func newError(kind Kind, ctx string) *Error {
	return &Error{
		kind: kind,
		ctx:  ctx,
	}
}

// newfError compiles and returns a new instance of the formatted Error.
func newfError(kind Kind, ctx string, args ...interface{}) *Error {
	return newError(kind, fmt.Sprintf(ctx, args...))
}

// wrapError compiles and returns a new instance of the Error with the
// context and underlying cause.
func wrapError(kind Kind, cause error, ctx string) *Error {
	return &Error{
		kind:  kind,
		ctx:   ctx,
		cause: cause,
	}
}

// wrapfError compiles and returns a new instance of the Error with the
// formatted context and underlying cause.
func wrapfError(kind Kind, cause error, ctx string, args ...interface{}) *Error {
	return wrapError(kind, cause, fmt.Sprintf(ctx, args...))
}
