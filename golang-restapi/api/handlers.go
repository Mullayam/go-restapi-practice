package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (app *application) Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World! %s", app.Domain)
}

func (app *application) Home(w http.ResponseWriter, r *http.Request) {
	var payload = struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
		Result  string `json:"result"`
	}{
		Success: true,
		Message: "Yayy",
		Result:  "1",
	}
	out, err := json.Marshal(payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(out)
}
