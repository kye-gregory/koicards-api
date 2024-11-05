package errors

import (
	"encoding/json"
	"log"
	"net/http"
)

type HttpErrorStack struct {
	StatusCode int					`json:"stausCode"`
	Errors     map[string][]string	`json:"errors,omitempty"`
}

func NewHttpErrorStack(status int) *HttpErrorStack {
	return &HttpErrorStack{StatusCode: status, Errors: make(map[string][]string)}
}

func (s *HttpErrorStack) Add(key string, err string) {
	s.Errors[key] = append(s.Errors[key], err)
}

func (s *HttpErrorStack) IsEmpty() bool {
	return (len(s.Errors) == 0)
}

func (s *HttpErrorStack) Error() string {
	if (len(s.Errors) == 0) { return "" }

	bytes, err := json.Marshal(s)
	if (err != nil) {
		log.Println(err.Error())
		return ""
	}
	return string(bytes)
}

// Returns either error or nil
func (s *HttpErrorStack) Return() *HttpErrorStack {
	if (len(s.Errors) > 0) { return s }
	return nil
}

// Return internal error
func (s *HttpErrorStack) ReturnInternalError() *HttpErrorStack {
	s.StatusCode = http.StatusInternalServerError

	// Clear Errors
	for k := range s.Errors {
		delete(s.Errors, k)
	}
	
	s.Errors["internal"] = []string{("internal server error")}
	return s
}