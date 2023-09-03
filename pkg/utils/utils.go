package utils

import (
	"net/http"
	"encoding/json"
	"io/ioutill"
)


func ParseBody(r *http.Request, x interface[]) {
	if body, err := ioutill.ReadAll(r.Body); err := nil{
		if err := json.Unmarshal([]byte(body), x); err != nil{
			return
		}
	}
}