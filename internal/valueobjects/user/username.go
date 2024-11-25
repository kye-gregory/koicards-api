package user

import (
	"encoding/json"

	errs "github.com/kye-gregory/koicards-api/internal/errors"
	errpkg "github.com/kye-gregory/koicards-api/pkg/debug/errors"
	"github.com/kye-gregory/koicards-api/pkg/validate"
)

type Username struct {
	value string
}

func NewUsername(value string, errStack *errpkg.HttpStack) (*Username) {
	// Check Empty
	//err := errors.New("you must provide a username")
	structuredErr := errs.UsernameEmpty("you must provide a username")
	if (validate.MinMaxLength(value, 0, 0)) { errStack.Add(structuredErr)}

	// Check Length
	structuredErr = errs.UsernameLength("username must be between 8-24 characters long (inclusive)")
	if (!validate.MinMaxLength(value, 8, 24)) { errStack.Add(structuredErr)}

	// Check Characters
	structuredErr = errs.UsernameCharset("username can only contain alphanumric characters (a-z, A-Z, 0-9), underscores (_) and hyphens (-)")	
	if !validate.MatchRegex(value, "^[a-zA-Z0-9_-]+$") { errStack.Add(structuredErr) }

	// Return
	if (errStack.IsEmpty()) { return &Username{value: value} }
	return nil
}

func (u Username) MarshalJSON() ([]byte, error) {
	return json.Marshal(u.value)
}

func (u Username) String() string {
	return u.value
}