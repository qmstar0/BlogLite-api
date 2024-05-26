package render

import (
	"errors"
	"github.com/qmstar0/domain/internal/pkg/e"
)

type Response struct {
	e.StateCode `json:",inline"`
	Data        any `json:"data,omitempty"`
}

func debugModeErrorResponse(err error) Response {
	var code e.StateCode
	if !errors.As(err, &code) {
		code = e.NotImplement.WithError(err)
	}
	return Response{
		StateCode: code,
	}
}

func releaseModeErrorResponse(err error) Response {
	var code e.StateCode
	if !errors.As(err, &code) {
		code = e.NotImplement.WithError(err)
	}
	code.Debug = ""
	return Response{
		StateCode: code,
	}
}
