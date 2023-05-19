package service

import (
	"encoding/json"
	"net/http"
)

// Constantes for header
const (
	HeaderContentType     = "Content-Type"
	HeaderApplicationJSON = "application/json"
)

// Payload is an interface for all type of data to return to client
type Payload interface {
	struct{} |
		GameService // Append here different data structures from database, separated by |
}

// Content is a structure containing the data and the description of the http response
type Content[P Payload] struct {
	Payload P      `json:"payload"`
	Details string `json:"details"`
}

// Encode methods sends content and ultimately sends error 500 in case of encoding error.
func (content *Content[any]) Encode(w http.ResponseWriter) {
	w.Header().Set(HeaderContentType, HeaderApplicationJSON)
	err := json.NewEncoder(w).Encode(content)
	if err != nil {
		// Last resort status (no content structure)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}
