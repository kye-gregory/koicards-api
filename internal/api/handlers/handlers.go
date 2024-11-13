package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	e "github.com/kye-gregory/koicards-api/pkg/errors"
)

func returnTextSuccess(w http.ResponseWriter, text string) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, text)
}


func returnHttpError(w http.ResponseWriter, stack *e.HttpErrorStack) bool {
	if !stack.IsEmpty() {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(stack.StatusCode)
		json.NewEncoder(w).Encode(stack)
		return true
	}

	return false
}