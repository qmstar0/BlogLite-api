package httperr

import (
	"common/e"
	"encoding/json"
	"fmt"
	"net/http"
)

const ok = `{"code":"OK"}`

type response struct {
	Code e.StateCode `json:"code"`
	Data any         `json:"data"`
}

func Respond(w http.ResponseWriter, data any) {
	respond(w, response{
		Code: e.Successed,
		Data: data,
	})
}

func Error(w http.ResponseWriter, err error) {
	respond(w, e.Unwrap(err))
}

func Success(w http.ResponseWriter) {
	_, err := fmt.Fprintf(w, ok)
	if err != nil {
		w.WriteHeader(502)
	}
}

func respond(w http.ResponseWriter, resp any) {
	marshal, err := json.Marshal(resp)
	if err != nil {
		return
	}
	_, err = w.Write(marshal)
	if err != nil {
		w.WriteHeader(502)
	}
}
