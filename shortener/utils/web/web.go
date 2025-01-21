package web

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

// Unmarshals Request's Body into out variable
func UnmarshalJSON(r *http.Request, out interface{}) error {
	// Check if request is empty
	if r.Body == nil {
		return errors.New("body is empty")
	}

	// Read Body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}

	// Unmarshal JSON
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
