package errorstack

import (
	"encoding/json"
	"log"
)

type ErrorStack interface {
	Error() string
	IsEmpty() bool
	Add(err error)
	Return() error
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