package utils

import (
	"encoding/json"
	"net/http"
)

func ParseBody(r *http.Request, x any) error {

	err := json.NewDecoder(r.Body).Decode(x)
	return err
}
