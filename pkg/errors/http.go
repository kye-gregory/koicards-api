package errors

import "net/http"

func ReturnHttpError(w http.ResponseWriter, err error, status int) bool {
	if err != nil {
		http.Error(w, err.Error(), status)
		return true
	}

	return false
}