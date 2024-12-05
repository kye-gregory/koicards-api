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
	// Append Error
	s.Errors = append(s.Errors, err)
}

func (s *HttpStack) Clear() {
	// Reset Error Slice
	s.Errors = make([]StructuredError, 0)
}

func (s *HttpStack) Contains(code ErrorCode) bool {
	// Loop Error Codes
	for _, err := range s.Errors {
		if err.Code() == code { return true }
	}

	// Return No-Find
	return false
}

func (s *HttpStack) IsEmpty() bool {
	// Check Length
	return (len(s.Errors) == 0)
}


func (s *HttpStack) Error() string {
	// Pre-Check For Errors
	if len(s.Errors) == 0 {
		return ""
	}

	// Marshal JSON
	bytes, err := json.Marshal(s)
	if err != nil {
		log.Println(err.Error())
		return ""
	}

	// Return Results
	return string(bytes)
}


func (s *HttpStack) InternalError(err StructuredError) {
	s.Clear()
	s.WithStatus(http.StatusInternalServerError)	
	s.Errors = []StructuredError{err}
}


func (s *HttpStack) WithStatus(status int) *HttpStack {
	s.StatusCode = status
	return s
}