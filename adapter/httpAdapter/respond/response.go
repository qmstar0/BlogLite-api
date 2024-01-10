package respond

import (
	"blog/apps/commandResult"
	"encoding/json"
	"net/http"
)

type responseTemplate struct {
	Code commandResult.StateCode
	Msg  string
}

func Respond(w http.ResponseWriter, code *commandResult.StateCode) {
	response := newResponseTemplate(*code)
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(response)
}

func newResponseTemplate(code commandResult.StateCode) *responseTemplate {
	msg := commandResult.MessageMap[code]
	return &responseTemplate{
		Code: code,
		Msg:  msg,
	}
}
