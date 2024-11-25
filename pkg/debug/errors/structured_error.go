package errors

import "encoding/json"

type ErrorField struct {
	value string
}

func (f ErrorField) MarshalJSON() ([]byte, error) {
	return json.Marshal(f.value)
}

func NewErrorField(value string) *ErrorField {
	ef := ErrorField{value: value}
	return &ef
}

type StructuredError struct {	
	field   ErrorField
	message string
	err     error
	code    string
}

func NewStructuredError(code, message string) *StructuredError {
	e := StructuredError{code: code, message: message}
	return &e
}

func (se StructuredError) WithError(err error) *StructuredError {
	se.err = err
	return &se
}

func (se StructuredError) WithField(field ErrorField) *StructuredError {
	se.field = field
	return &se
}

func (se StructuredError) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Field   ErrorField `json:"field,omitempty"`
		Message string     `json:"message"`
		Err     error      `json:"err,omitempty"`
		Code    string     `json:"code"`
	}{
		Field: se.field,
		Message: se.message,
		Err: se.err,
		Code: se.code,
	})
}