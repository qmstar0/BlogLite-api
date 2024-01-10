package respond

import (
	"blog/apps/commandResult"
	"fmt"
	"strings"
)

func NewError(code commandResult.StateCode, err ...error) error {
	resp := newError(code)
	resp.Warp(err...)
	return resp
}

func newError(code commandResult.StateCode) *responseTemplate {
	msg := commandResult.MessageMap[code]
	return &responseTemplate{
		Code: code,
		Msgs: []string{msg},
	}
}

type responseTemplate struct {
	Code commandResult.StateCode
	Msgs []string
}

func (r *responseTemplate) Error() string {
	return fmt.Sprintf("%d:%s", r.Code, r.Msgs)
}

func (r *responseTemplate) Warp(errs ...error) {
	for _, err := range errs {
		r.Msgs = append(r.Msgs, err.Error())
	}
}

func (r *responseTemplate) toMap() map[string]any {
	result := make(map[string]any)
	result["code"] = r.Code
	result["msg"] = strings.Join(r.Msgs, "; ")
	return result
}
