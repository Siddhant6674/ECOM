package utils

import (
	"encoding/json"
	"fmt"

	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
)

var Validate = validator.New()

func ParseJSON(r *http.Request, payload any) error {
	if r.Body != nil {
		return fmt.Errorf("missing request body")
	}
	return json.NewDecoder(r.Body).Decode(payload)
}

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

// func WriteError(w http.ResponseWriter, status int, err error) {
// 	// WriteJSON(w, status, map[string]string{"error": err.Error()})

// }

func WriteError(w http.ResponseWriter, status int, err error) {
	log.Printf("error -> %s", err.Error())
	w.WriteHeader(status)
	w.Write([]byte(""))
}
