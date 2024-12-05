package errors

import "encoding/json"

/*/ ERROR FIELD /*/
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



/*/ ERROR CODE /*/
type ErrorCode struct {
	value string
}

func (c ErrorCode) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func NewErrorCode(value string) *ErrorCode {
	ec := ErrorCode{value: value}
	return &ec
}



/*/ STRUCTURED ERROR /*/
type StructuredError struct {	
	field   ErrorField
	message string
	err     error
	code    ErrorCode
}

func NewStructuredError(code ErrorCode, message string) *StructuredError {
	e := StructuredError{code: code, message: message}
	return &e
}

func (se *StructuredError) WithError(err error) *StructuredError {
	se.err = err
	return se
}

func (se *StructuredError) WithField(field ErrorField) *StructuredError {
	se.field = field
	return se
}

func (se StructuredError) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Field   string `json:"field,omitempty"`
		Message string     `json:"message"`
		Err     error      `json:"err,omitempty"`
		Code    string  `json:"code"`
	}{
		Field: se.field.value,
		Message: se.message,
		Err: se.err,
		Code: se.code.value,
	})
}

func (se StructuredError) Code() ErrorCode {
	return se.code
}