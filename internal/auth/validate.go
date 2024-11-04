package auth

import "errors"

func ValidateEmail(s string) error {
	return nil
}

func ValidateUsername(s string) error {
	if (len(s) < 8) {return errors.New("username length must be more than 8 characters")}
	return nil
}

func ValidatePassword(s string) error {
	if (len(s) < 8) {return errors.New("password length must be more than 8 characters")}
	return nil
}