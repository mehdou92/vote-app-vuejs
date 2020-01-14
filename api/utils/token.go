package utils

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"

	"github.com/lavaninho/Projet-GO/models"
)

// CheckJwtToken check if jwt token is valid
func CheckJwtToken(tokenString string) (jwt.MapClaims, error) {
	signingKey := []byte(os.Getenv("JWT_SECRET"))

	// check jwt token is exist.
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return signingKey, nil
	})

	if err != nil {
		return nil, err
	}

	claims, _ := token.Claims.(jwt.MapClaims)

	return claims, err
}

// GenerateJWT generate jwt token
func GenerateJWT(u models.User) (string, error) {
	signingKey := []byte(os.Getenv("JWT_SECRET"))

	// set token expr time, username, email and id to token.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":          time.Now().Add(time.Hour * 12 * 1).Unix(),
		"uuid":         u.UUID,
		"email":        u.Email,
		"access_level": u.AccessLevel,
	})

	tokenString, err := token.SignedString(signingKey)

	return tokenString, err
}

// CheckPassword check user password
func CheckPassword(u models.User, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))

	return err == nil
}
