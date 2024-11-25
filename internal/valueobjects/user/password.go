package user

import (
	"errors"
	"unicode/utf8"

	"github.com/kye-gregory/koicards-api/internal/auth"
	errpkg "github.com/kye-gregory/koicards-api/pkg/debug/errors"
)

type Password struct {
	hashed string
}

func NewPassword(value string, errStack *errpkg.HttpStack) *Password {
	// Check Password Isn't Empty
	err := errors.New("you must provide a password")
	if utf8.RuneCountInString(value) == 0 { errStack.Add("password", err) }

	// Check Password Length
	err = errors.New("password length must be between 8-64 characters (inclusive)")
	if utf8.RuneCountInString(value) < 8 { errStack.Add("password", err) }
	if utf8.RuneCountInString(value) > 64 { errStack.Add("password", err) }

	// TODO: Add More Password Validations!
	//	- Character Variety (i.e at least one number, capital, and special character)?
	//	- Prevent Common Patterns (i.e password)
	//	- No username in password

	// Hash Password
	hashed, err := auth.Hash(value)
	if (err != nil) { errStack.ReturnInternalError() }

	// Return
	if (errStack.IsEmpty()) { return &Password{hashed: hashed} }
	return nil
}

func (p Password) String() string {
	return p.hashed
}