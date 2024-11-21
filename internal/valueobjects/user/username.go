package user

import (
	"encoding/json"
	"errors"

	"github.com/kye-gregory/koicards-api/pkg/debug/errorstack"
	"github.com/kye-gregory/koicards-api/pkg/validate"
)

type Username struct {
	value string
}

func NewUsername(value string, errStack *errorstack.HttpStack) (*Username) {
	// Check Empty
	err := errors.New("you must provide a username")
	if (validate.MinMaxLength(value, 0, 0)) { errStack.Add("username", err)}

	// Check Length
	err = errors.New("username length must be between 8-24 characters (inclusive)")
	if (!validate.MinMaxLength(value, 8, 24)) { errStack.Add("username", err)}

	// Check Characters
	err = errors.New("username can only contain alphanumric characters (a-z, A-Z, 0-9), underscores (_) and hyphens (-)")	
	if !validate.MatchRegex(value, "^[a-zA-Z0-9_-]+$") { errStack.Add("username", err) }

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