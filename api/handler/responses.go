package handler

import (
	"encoding/json"
	"net/http"

	"github.com/sotnikov-s/glossarize-backend/api/errors"
)

// JSONResponse is a type that represents
// a server response in the JSON format.
type JSONResponse struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data,omitempty"`
	Error  interface{} `json:"error,omitempty"`
}

// WriteSuccessResponse writes a success json response.
func WriteSuccessResponse(w http.ResponseWriter, status int, data interface{}) error {
	return WriteJSON(w, status, JSONResponse{Status: status, Data: data})
}

// WriteErrorResponse writes an error json response.
func WriteErrorResponse(w http.ResponseWriter, status int, err interface{}) error {
	return WriteJSON(w, status, JSONResponse{Status: status, Error: err})
}

// WriteJSON writes a json response.
func WriteJSON(w http.ResponseWriter, status int, body interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(body); err != nil {
		return errors.Internal.Wrap(err, "failed to write the response body")
	}
	return nil
}
