package errors

import (
	"encoding/json"
	"log"
	"net/http"
)

type HttpStack struct {
	StatusCode int                  `json:"statusCode"`
	Errors     []StructuredError	`json:"errors,omitempty"`
}


func NewHttpStack() *HttpStack {
	return &HttpStack{StatusCode: http.StatusInternalServerError, Errors: make([]StructuredError, 0)}
}

func (s *HttpStack) Add(err StructuredError) {
	s.Errors = append(s.Errors, err)
}

func (s *HttpStack) Clear() {
	s.Errors = make([]StructuredError, 0)
}

func (s *HttpStack) Contains(code ErrorCode) bool {
	for _, err := range s.Errors {
		if err.Code() == code { return true }
	}

	return false
}

func (s *HttpStack) IsEmpty() bool {
	return (len(s.Errors) == 0)
}

// Returns errors as JSON
func (s *HttpStack) Error() string {
	if len(s.Errors) == 0 {
		return ""
	}

	bytes, err := json.Marshal(s)
	if err != nil {
		log.Println(err.Error())
		return ""
	}

	return string(bytes)
}

// Return internal error
func (s *HttpStack) InternalError(err StructuredError) {
	s.Clear()
	s.WithStatus(http.StatusInternalServerError)	
	s.Errors = []StructuredError{err}
}

// Returns either error or nil
func (s *HttpStack) Return() error {
	if len(s.Errors) > 0 { return s }
	return nil
}

// Update Status
func (s *HttpStack) WithStatus(status int) *HttpStack {
	s.StatusCode = status
	return s
}