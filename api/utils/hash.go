package utils

import "golang.org/x/crypto/bcrypt"

// Hash hash password
func Hash(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 4)
	return string(bytes)
}
