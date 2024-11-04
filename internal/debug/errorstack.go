package debug

import (
	"fmt"
	"time"
)

// ErrorStack collects multiple errors and implements the error interface.
type ErrorStack struct {
	Errors []error
}

// Appends an error to the ErrorStack.
func (m *ErrorStack) Add(err error) {
	if err != nil {
		timestampedErr := fmt.Errorf("%s: %w", time.Now().Format(time.RFC3339), err)
		m.Errors = append(m.Errors, timestampedErr)
	}
}

// Implements the error interface for ErrorStack.
func (m *ErrorStack) Error() string {
	if len(m.Errors) == 0 {
		return "no errors"
	}
	msg := "multiple errors:"
	for _, err := range m.Errors {
		msg += "\n - " + err.Error()
	}
	return msg
}