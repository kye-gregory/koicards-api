package user

import (
	"encoding/json"
	"errors"
	"regexp"
	"unicode/utf8"

	"github.com/kye-gregory/koicards-api/pkg/debug/errorstack"
)

type Username struct {
	value string
}

func NewUsername(value string, errStack *errorstack.HttpStack) (*Username) {
	// Check Username Isn't Empty
	err := errors.New("you must provide a username")
	if utf8.RuneCountInString(value) == 0 { errStack.Add("username", err) }

	// Check Username Length
	err = errors.New("username length must be between 8-24 characters (inclusive)")
	if utf8.RuneCountInString(value) < 8 { errStack.Add("username", err) }
	if utf8.RuneCountInString(value) > 24 { errStack.Add("username", err) }

	// Check Username Characters
	err = errors.New("username can only contain alphanumric characters (a-z, A-Z, 0-9), underscores (_) and hyphens (-)")
	match, _ := regexp.MatchString("^[a-zA-Z0-9_-]+$", value)
	if !match { errStack.Add("username", err) }

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