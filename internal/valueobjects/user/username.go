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
	// Check Length
	structuredErr := errs.UsernameLength("username must be between 8-24 characters long (inclusive)")
	if !validate.MinMaxLength(value, 8, 24) { errStack.Add(structuredErr)}

	// Check Characters
	structuredErr = errs.UsernameCharset("username can only contain alphanumeric characters (a-z, A-Z, 0-9), underscores (_) and hyphens (-)")	
	if !validate.MatchRegex(value, "^[\\w_-]+$") { errStack.Add(structuredErr) }

	// Check Format
	structuredErr = errs.UsernameFormat("username cannot have more than one allowed special character in a row")	
	if validate.MatchRegex(value, ".*[_-]{2,}.*") { errStack.Add(structuredErr) }

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