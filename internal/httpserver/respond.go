package httpserver

import (
	"encoding/json"
	"net/http"
)

// Success is the response format when http handler succeeds
type Success struct {
	Message string `json:"message,omitempty"`
}

func RespondJSON(w http.ResponseWriter, obj interface{}) {
	w.Header().Set("Content-Type", "application/json")

	status := http.StatusOK
	var respBytes []byte
	var err error

	switch parsed := obj.(type) {
	case *Error:
		if parsed.Status >= http.StatusInternalServerError && parsed.Status != http.StatusServiceUnavailable {
			parsed.Desc = DefaultErrorDesc
		}
		status = parsed.Status
		if status == 0 {
			status = http.StatusInternalServerError
		}
		respBytes, err = json.Marshal(parsed)
	case error:
		status = http.StatusInternalServerError
		respBytes, err = json.Marshal(ErrDefaultInternal)
	default:
		respBytes, err = json.Marshal(obj)
	}

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Write response
	w.WriteHeader(status)
	w.Write(respBytes)
}
