package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
)

var Validate = validator.New()

func GetTokenFromRequest(r *http.Request) (string, error) {
	tokenString := strings.TrimPrefix(r.Header.Get("Authorization"), "Bearer ")
	if tokenString == "" {
		return "", errors.New("missing token in header")
	}
	return tokenString, nil
}

func ParseJSON(r *http.Request, payload interface{}) error {
	if r.Body == nil {
		return fmt.Errorf("missing request body")
	}
	return json.NewDecoder(r.Body).Decode(payload)
}

func RespondWithJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(v)
}

func RespondWithError(w http.ResponseWriter, status int, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	_ = json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
}
