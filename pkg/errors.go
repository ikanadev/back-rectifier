package pkg

import (
	"errors"
	"fmt"
)

// APIError represents a controlled API user error, everything else will be a 500 error
type APIError struct {
	Message string `json:"message"`
	Code    string `json:"code"`
	Err     error  `json:"error"`
}

// NewAPIErr creates a new error
func NewAPIErr(message, code string) APIError {
	return APIError{
		Message: message,
		Code:    code,
		Err:     errors.New(message),
	}
}

func (e APIError) Error() string {
	return fmt.Sprintf("Message: %s \n Error: %s", e.Message, e.Err.Error())
}

// IsAPIError returns true is the current error is an espected error
func IsAPIError(e error) bool {
	_, ok := e.(APIError)
	return ok
}
