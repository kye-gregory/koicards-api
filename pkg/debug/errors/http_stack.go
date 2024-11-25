package errors

import (
	"encoding/json"
	"log"
	"net/http"
)

type HttpStack struct {
	StatusCode int                 `json:"statusCode"`
	Errors     map[string][]string `json:"errors,omitempty"`
}


func NewHttpStack() *HttpStack {
	return &HttpStack{StatusCode: http.StatusInternalServerError, Errors: make(map[string][]string)}
}


func (s *HttpStack) Status(status int) *HttpStack {
	s.StatusCode = status
	return s
}

func (s *HttpStack) Clear() *HttpStack {
	for k := range s.Errors {
		delete(s.Errors, k)
	}
	return s
}


func (s *HttpStack) Add(key string, err error) {
	s.Errors[key] = append(s.Errors[key], err.Error())
}


func (s *HttpStack) IsEmpty() bool {
	return (len(s.Errors) == 0)
}


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


// Returns either error or nil
func (s *HttpStack) Return() error {
	if len(s.Errors) > 0 { return s }
	return nil
}


// Return internal error
func (s *HttpStack) ReturnInternalError() {
	s.Clear().Status(http.StatusInternalServerError)	
	s.Errors["internal"] = []string{"internal server error"}
}