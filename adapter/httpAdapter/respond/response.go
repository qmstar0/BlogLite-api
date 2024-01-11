package respond

import (
	"blog/apps/commandResult"
	"encoding/json"
	"errors"
	"net/http"
)

func Respond(w http.ResponseWriter, err *error) {
	var respE *responseTemplate
	if !errors.As(*err, &respE) {
		respE = newError(commandResult.NotImplementedErr)
	}
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(respE.toResponse())
}
