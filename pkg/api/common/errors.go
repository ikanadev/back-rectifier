package common

import "errors"

var (
	// ErrNoContextID when id cant be readed from context
	ErrNoContextID = errors.New("Cannot read ID in context (missing token?)")
	// ErrContextIDInvalid when the context id can't be converted to int
	ErrContextIDInvalid = errors.New("Context ID can't be converted to int")
)
