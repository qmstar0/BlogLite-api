package bind

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

var ParameterParsingError = errors.New("参数解析错误")

func Decode(r *http.Request, v any) error {
	return decodeJSON(r.Body, v)
}

func decodeJSON(body io.ReadCloser, v any) error {
	return json.NewDecoder(body).Decode(v)
}
