package handlers

import "net/http"

func ReturnHttpError(w http.ResponseWriter, err error, code int) bool {
	if err != nil {
		status := http.StatusBadRequest
		http.Error(w, err.Error(), status)
		return true
	}

	return false
}