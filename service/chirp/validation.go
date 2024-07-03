package chirp

import (
	"errors"
	"strings"
)

var (
	maxChirpLengsth = 140
	badWords        = map[string]*string{
		"kerfuffle": nil,
		"sharbert":  nil,
		"fornax":    nil,
	}
)

func ValidateChirp(msg string) (string, error) {
	if len(msg) >= maxChirpLengsth {
		return "", errors.New("More than 140 characters")
	}

	words := strings.Split(msg, " ")
	for i, word := range words {
		if _, ok := badWords[strings.ToLower(word)]; ok {
			words[i] = "****"
		}
	}
	return strings.Join(words, " "), nil

}
