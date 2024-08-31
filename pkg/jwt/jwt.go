package jwt

import (
	"crypto/rand"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"log"

	"time"
)

var secretKet = generateSecretKey(32)
var refreshSecretKey = generateSecretKey(32)

func CreateToken(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Minute * 15).Unix(),
		"iat":      time.Now().Unix(),
		"app":      "note-service-token",
	})

	tokenString, err := token.SignedString(secretKet)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKet, nil
	})
	if err != nil {
		return err
	}
	if !token.Valid {
		return fmt.Errorf("invalid token")
	}

	return nil
}

func CreateRefreshToken(username string) (string, error) {
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 24 * 7).Unix(), // Срок действия 7 дней
		"iat":      time.Now().Unix(),
		"app":      "note-service-refresh-token",
	})

	refreshTokenString, err := refreshToken.SignedString(refreshSecretKey)
	if err != nil {
		return "", err
	}

	return refreshTokenString, nil
}

func VerifyRefreshToken(refreshTokenString string) error {
	token, err := jwt.Parse(refreshTokenString, func(token *jwt.Token) (interface{}, error) {
		return refreshSecretKey, nil
	})
	if err != nil {
		return err
	}
	if !token.Valid {
		return fmt.Errorf("invalid refresh token")
	}

	return nil
}

func generateSecretKey(length int) []byte {
	key := make([]byte, length)
	_, err := rand.Read(key)
	if err != nil {
		log.Fatal(err)
	}
	return key
}
