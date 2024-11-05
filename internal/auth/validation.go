package auth

import (
	"errors"
	"net/mail"
	"regexp"

	e "github.com/kye-gregory/koicards-api/pkg/errors"
)


func ValidateEmail(errStack *e.HttpErrorStack, email string) {
	// Parse Email
	err := errors.New("you must provide a valid email (i.e johndoe@example.com)")
	_, parseErr := mail.ParseAddress(email)
	if (parseErr != nil) { errStack.Add("email", err.Error()) }
}


func ValidateUsername(errStack *e.HttpErrorStack, username string) {
	// Check Username Isn't Empty
	err := errors.New("you must provide a username")
	if (len(username) == 0) { errStack.Add("username", err.Error()) }

	// Check Username Length
	err = errors.New("username length must be between 8-24 characters (inclusive)")
	if (len(username) < 8) { errStack.Add("username", err.Error()) }
	if (len(username) > 24) { errStack.Add("username", err.Error()) }

	// Check Username Characters
	err = errors.New("username can only contain alphanumric characters (a-z, A-Z, 0-9), underscores (_) and hyphens (-)")
	match, _ := regexp.MatchString("^[a-zA-Z0-9_-]+$", username)
	if (!match) { errStack.Add("username", err.Error()) }
}


func ValidatePassword(errStack *e.HttpErrorStack, password string) {
	// Check Password Isn't Empty
	err := errors.New("you must provide a password")
	if (len(password) == 0) { errStack.Add("password", err.Error()) }

	// Check Password Length
	err = errors.New("password length must be between 8-64 characters (inclusive)")
	if (len(password) < 8) { errStack.Add("password", err.Error()) }
	if (len(password) > 64) { errStack.Add("password", err.Error()) }

	// TODO: Add More Password Validations!
	//	- Character Variety (i.e at least one number, capital, and special character)?
	//	- Prevent Common Patterns (i.e password)
	//	- No username in password
}