package user

import (
	"encoding/json"
	"errors"
	"net/mail"

	"github.com/kye-gregory/koicards-api/pkg/debug/errorstack"
)

type Email struct {
	value string
}

func NewEmail(value string, errStack *errorstack.HttpStack) (*Email) {

	// Parse Email
	err := errors.New("you must provide a valid email (i.e johndoe@example.com)")
	_, parseErr := mail.ParseAddress(value)
	if (parseErr != nil) { errStack.Add("email", err) }

	// Return
	if (errStack.IsEmpty()) { return &Email{value: value} }
	return nil
}

func (e Email) MarshalJSON() ([]byte, error) {
	return json.Marshal(e.value)
}

func (e Email) String() string {
	return e.value
}