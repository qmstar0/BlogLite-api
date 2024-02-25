package httperr

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type response struct {
	Code    StateCode `json:"code"`
	Message string    `json:"message"`
}

func Respond(w http.ResponseWriter, code StateCode, info string) {
	reply(w, response{
		Code:    code,
		Message: fmt.Sprintf("%s;%s", code.Error(), info),
	})
}

func reply(w http.ResponseWriter, resp response) {
	data, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func Success(w http.ResponseWriter) {
	reply(w, response{
		Code:    Successed,
		Message: Successed.Error(),
	})
}
