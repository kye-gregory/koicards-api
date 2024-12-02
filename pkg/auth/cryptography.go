package auth

import "golang.org/x/crypto/bcrypt"

func Hash(s string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(s), 10)
	return string(bytes), err
}

func VerifyPassword(hash, input string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(input))
	return err == nil // Returns true if the password matches
}