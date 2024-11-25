package user

import (
	"encoding/json"
	"net/mail"

	errs "github.com/kye-gregory/koicards-api/internal/errors"
	errpkg "github.com/kye-gregory/koicards-api/pkg/debug/errors"
)

type Email struct {
	value string
}

func NewEmail(value string, errStack *errpkg.HttpStack) (*Email) {
	// Validate Email
	structuredErr := errs.EmailInvalid("you must provide a valid email (i.e johndoe@example.com)")
	_, err := mail.ParseAddress(value)
	if (err != nil) { errStack.Add(structuredErr) }

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