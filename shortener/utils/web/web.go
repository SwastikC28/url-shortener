package web

import (
	"encoding/json"

	"io"
	"net/http"

	"url-shortener/utils/errors"
)

func UnmarshalJSON(r *http.Request, out interface{}) error {
	if r.Body == nil {
		return errors.NewValidationError("Body cannot be empty")
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

func RespondError(w http.ResponseWriter, err error) {
	w.Header().Set("content-type", "application/json")

	switch err.(type) {
	case *errors.ValidationError:
		w.WriteHeader(http.StatusBadRequest)
	case *errors.ResourceNotFoundError:
		w.WriteHeader(http.StatusNotFound)
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Write([]byte(err.Error()))
}
