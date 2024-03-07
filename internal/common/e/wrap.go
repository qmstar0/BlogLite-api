package e

import (
	"errors"
	"fmt"
)

type StateError struct {
	Code    StateCode `json:"code"`
	Message string    `json:"message"`
}

func (e StateError) Error() string {
	return e.Message
}

func Wrap(code StateCode, err error) error {
	var stateErr StateError
	if errors.As(err, &stateErr) {

		stateErr.Message = fmt.Sprintf("%s; %s", code.Error(), stateErr.Message)
		return stateErr
	}

	return StateError{
		Code:    code,
		Message: code.Error(),
	}
}

func Unwrap(err error) StateError {
	var stateErr StateError
	if errors.As(err, &stateErr) {
		return stateErr
	}
	return StateError{
		Code:    NotImplementedErr,
		Message: err.Error(),
	}
}
