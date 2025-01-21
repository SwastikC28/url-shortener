package web

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

func UnmarshalJSON(r *http.Request, out interface{}) error {
	if r.Body == nil {
		return errors.New("body is empty")
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, out)
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
