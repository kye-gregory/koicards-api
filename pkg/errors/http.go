package errors

import (
	"encoding/json"
	"net/http"
)

func ReturnHttpError(w http.ResponseWriter, stack *HttpErrorStack) bool {
	if !stack.IsEmpty() {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(stack.StatusCode)
		json.NewEncoder(w).Encode(stack)
		return true
	}

	return false
}