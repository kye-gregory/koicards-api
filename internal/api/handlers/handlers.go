package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/kye-gregory/koicards-api/pkg/debug/errorstack"
)

func returnSuccess(w http.ResponseWriter, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(payload)
}


func returnHttpError(w http.ResponseWriter, stack *errorstack.HttpStack) bool {
	if !stack.IsEmpty() {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(stack.StatusCode)
		json.NewEncoder(w).Encode(stack)
		return true
	}

	return false
}