package web

import (
	"encoding/json"
	"net/http"
)

func UnmarshalJSON(r *http.Request, data interface{}) error {
	body, err := json.Marshal(r.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, data)
	if err != nil {
		return err
	}

	return nil
}

func RespondJSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("content-type", "application/json")

	body, err := json.Marshal(data)
	if err != nil {
		return
	}

	w.WriteHeader(statusCode)
	w.Write(body)
}
