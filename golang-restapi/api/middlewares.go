package main

import "net/http"

func (app *application) enableCors(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		if r.Method == "OPTIONS" {
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			w.Header().Set("Access-Control-Allow-Headers", "Authorization,Access-Control-Allow-Headers, Origin,Accept, X-Requested-With, Content-Type, Access-Control-Request-Method, Access-Control-Request-Headers")
			return
		} else {
			h.ServeHTTP(w, r)
		}

	})
}
func (app *application) isAuthenticated(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _, err := app.auth.GetTokenFromHeaderAndVerify(w, r)
		if err != nil {
			app.errorJSON(w, err, http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)

	})
}
