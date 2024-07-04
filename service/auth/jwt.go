package auth

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/mathiaskluge/chirpx/config"
)

func CreateJWT(secret string, userID int, expiresInSeconds int) (string, error) {
	expiration := time.Second * time.Duration(expiresInSeconds)

	if expiresInSeconds == 0 || expiresInSeconds >= config.Env.JWTExpirationInSeconds {
		expiration = time.Second * time.Duration(config.Env.JWTExpirationInSeconds)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID": strconv.Itoa(userID),
		"iat":    time.Now().Unix(),
		"exp":    time.Now().Add(expiration).Unix(),
	})

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateJWT(tokenString string, secret string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("token invalid or expired")
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		return claims, nil
	} else {
		return nil, errors.New("unable to parse claims")
	}
}
