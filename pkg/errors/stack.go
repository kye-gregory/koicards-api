package errors

// Stack collects multiple errors and implements the error interface.
type Stack struct {
	Errors []error
}

// Creates a new Stack
func NewStack() *Stack {
	s := new(Stack)
	return s
}

// Appends an error to the Stack.
func (s *Stack) Add(err error) {
	if err != nil {
		s.Errors = append(s.Errors, err)
	}
}

// Implements the error interface.
func (s *Stack) Error() string {
	// No Errors
	if len(s.Errors) == 0 { return "" }

	// Construct Err String
	msg := "\nErrors:"
	for _, err := range s.Errors {
		msg += "\n - " + err.Error()
	}

	// Return
	return msg
}

// Resets the Stack
func (s *Stack) Clear() {
	s.Errors = nil
}

// Returns either error or nil
func (s *Stack) Return() error {
	if (len(s.Errors) > 0) { return s }
	return nil
}