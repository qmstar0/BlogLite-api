package binder

import (
	"encoding/json"
	"io"
	"net/http"
)

type Binder interface {
	Bind(r *http.Request, i any) error
}
type JSONBinder struct {
}

func (J JSONBinder) Bind(r *http.Request, i any) error {
	all, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(all, i)
}
