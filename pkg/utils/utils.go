package utils

import (
	"encoding/json"
	"io"
	"net/http"
)

// we need unmarshal data, since it will be in json (data from user)

func ParseBody(r *http.Request, x interface{}) {
	if body, err := io.ReadAll(r.Body); err == nil {
		if err := json.Unmarshal([]byte(body), x); err != nil {
			return
		}
	}
}
