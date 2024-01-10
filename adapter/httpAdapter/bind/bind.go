package bind

import (
	"encoding/json"
	"io"
	"net/http"
)

func Decode(r *http.Request, v any) error {
	return decodeJSON(r.Body, v)
}

func decodeJSON(body io.ReadCloser, v any) error {
	return json.NewDecoder(body).Decode(v)
}
