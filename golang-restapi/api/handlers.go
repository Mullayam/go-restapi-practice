package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
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
	// app.DB.GetAll()
	app.writeJSON(w, http.StatusOK, payload)
	out, err := json.Marshal(payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(out)
}
func (app *application) authenticate(w http.ResponseWriter, r *http.Request) {
	// app.DB.GetAll()
	var reqPayload struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	err := app.readJSON(w, r, &reqPayload)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	u := JwtUser{1, "John", "Doe"}
	token, err := app.auth.GenerateTokenPair(&u)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	w.Write([]byte(token.Token))
}
func (j *Auth) GetRefreshcookie(refreshToken string) *http.Cookie {
	return &http.Cookie{
		Name:     j.CookieName,
		Path:     j.CookiePath,
		Value:    refreshToken,
		Expires:  time.Now().Add(j.RefreshExpiry),
		MaxAge:   int(j.RefreshExpiry.Seconds()),
		Secure:   false,
		Domain:   j.CookieDomain,
		HttpOnly: true,
	}
}
func (j *Auth) GetExpiredRefreshcookie() *http.Cookie {
	return &http.Cookie{
		Name:     j.CookieName,
		Path:     j.CookiePath,
		Value:    "",
		Expires:  time.Unix(0, 0),
		MaxAge:   -1,
		Secure:   false,
		Domain:   j.CookieDomain,
		HttpOnly: true,
	}
}
