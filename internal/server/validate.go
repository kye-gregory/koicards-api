package server

import "golang.org/x/crypto/bcrypt"

func hash(s string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(s), 10)
	return string(bytes), err
}