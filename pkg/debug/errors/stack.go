package errors

import (
	"encoding/json"
	"log"
)

type ErrorStack interface {
	Error() string
	IsEmpty() bool
	Add(err error)
	Return() error
	Clear()
}

type Stack struct {
	Errors map[string][]string `json:"errors,omitempty"`
}

// Creates a new Stack
func NewStack() *Stack {
	s := new(Stack)
	return s
}

// Appends an error to the Stack.
func (s *Stack) Add(key string, err error) {
	s.Errors[key] = append(s.Errors[key], err.Error())
}

// Clears all errors on stack
func (s *Stack) Clear() *Stack {
	for k := range s.Errors {
		delete(s.Errors, k)
	}
	return s
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

// Returns either error or nil
func (s *Stack) Return() error {
	if len(s.Errors) > 0 {
		return s
	}
	return nil
}