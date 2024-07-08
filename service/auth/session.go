package auth

import (
	"crypto/rand"
	"encoding/hex"
)

// Creates a 32-byte (256-bit) long unique session IDs using rand.Read()
// Returns it and saves it to the database.
// with P(n) approx. 1 - e^(-4.32*10^-42) probabilyt for a collision
// and should therefore be reliable for this use case.
//
// Ensuring uniqueness is important to enable storing IDs as map keys for
// O(1) lookups in most cases.
func GenerateSessionID() (string, error) {
	id := make([]byte, 32)

	_, err := rand.Read(id)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(id), nil
}
