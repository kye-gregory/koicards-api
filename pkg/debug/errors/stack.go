package errors

import (
	"encoding/json"
	"log"
)

type ErrorStack interface {
	Add(err StructuredError)
	Clear()
	Contains(code ErrorCode) bool
	IsEmpty() bool
	Error() string
	InternalError(err StructuredError)
	Return() error
}

type Stack struct {
	Errors []StructuredError `json:"errors,omitempty"`
}

// Creates a new Stack
func NewStack() *Stack {
	s := new(Stack)
	return s
}

func (s *Stack) Add(err StructuredError) {
	s.Errors = append(s.Errors, err)
}

func (s *Stack) Clear() {
	s.Errors = make([]StructuredError, 0)
}

func (s *Stack) Contains(code ErrorCode) bool {
	for _, err := range s.Errors {
		if err.Code() == code { return true }
	}

	return false
}

func (s *Stack) IsEmpty() bool {
	return (len(s.Errors) == 0)
}

// Implements the error interface.
func (s *Stack) Error() string {
	// No Errors
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

func (s *Stack) InternalError(err StructuredError) {
	s.Clear()
	s.Errors = []StructuredError{err}
}

// Returns either error or nil
func (s *Stack) Return() error {
	if len(s.Errors) > 0 {
		return s
	}
	return nil
}