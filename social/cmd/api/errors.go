package main

import "net/http"

func (app *application) internalServer(w http.ResponseWriter, message string) {
	writeJSONError(w, http.StatusInternalServerError, message)
}

// bad requiest error
// validation error
// not found error
// forbidden error
// unauthorized error
