package user

import (
	errs "github.com/kye-gregory/koicards-api/internal/errors"
	"github.com/kye-gregory/koicards-api/pkg/auth"
	errpkg "github.com/kye-gregory/koicards-api/pkg/debug/errors"
	"github.com/kye-gregory/koicards-api/pkg/validate"
)

type Password struct {
	hashed string
}

func NewPassword(value string, errStack *errpkg.HttpStack) *Password {
	// Check Password Length
	structuredErr := errs.PasswordLength("password must be between 8-64 characters long (inclusive)")
	if (!validate.MinMaxLength(value, 8, 64)) { errStack.Add(structuredErr)}

	// TODO: Add More Password Validations!
	//	- Character Variety (i.e at least one number, capital, and special character)?
	//	- Prevent Common Patterns (i.e password)
	//	- No username in password

	// Hash Password
	hashed, err := auth.Hash(value)
	if err != nil { errs.Internal(errStack, err); return nil }

	// Return
	if (errStack.IsEmpty()) { return &Password{hashed: hashed} }
	return nil
}

func NewPasswordFromDB(value string) *Password {
	return &Password {hashed: value}
}

func (p Password) String() string {
	return p.hashed
}