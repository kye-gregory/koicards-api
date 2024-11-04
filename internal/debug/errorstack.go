package debug

import (
	"fmt"
	"log"
	"time"
)

// ErrorStack collects multiple errors and implements the error interface.
type ErrorStack struct {
	Errors []error
}

// Appends an error to the ErrorStack.
func (m *ErrorStack) Add(err error) {
	if err != nil {
		log.Println("ERROR", err)
		timestampedErr := fmt.Errorf("%s ERROR %w", time.Now().Format(TIME_FORMAT), err)
		m.Errors = append(m.Errors, timestampedErr)
	}
}

// Implements the error interface for ErrorStack.
func (m *ErrorStack) Error() string {
	if len(m.Errors) == 0 {
		return ""
	}
	msg := "\nErrors:"
	for _, err := range m.Errors {
		msg += "\n - " + err.Error()
	}
	return msg + "\n"
}