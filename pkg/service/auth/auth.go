package auth

import (
	"errors"
	"os"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CompareHashPassword(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

type JwtClaims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func ParseToken(token string) (JwtClaims, error) {
	claims := JwtClaims{}

	res, err := jwt.ParseWithClaims(token, &claims, func(t *jwt.Token) (interface{}, error) {
		return os.ReadFile("./public.key")
	}, jwt.WithValidMethods([]string{"RS256","RS384","RS512"}))
	if err != nil {
		return claims, err
	}

	if res.Valid {
		return claims, nil
	}
	return claims, errors.New("Unauthorized")
}
