package auth

import "golang.org/x/crypto/bcrypt"

func HashPassword(pw string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func ComparePasswords(pwHash string, plain []byte) bool {
	err := bcrypt.CompareHashAndPassword([]byte(pwHash), plain)
	return err == nil
}
