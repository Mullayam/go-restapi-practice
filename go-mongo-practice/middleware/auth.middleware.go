package middleware

import "net/http"

func JwtAuthMiddleware(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) http.Handler {
	authHeader := r.Header.Get("Authorization")

	//validate token
	if authHeader != "Bearer token" {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Unauthorized"))

	}
	//if token is valid, call next(w, r)
	return next(w, r)

}

func AdminMiddleware() {}

func UserMiddleware() {}
func EnableCors(h http.Handler) http.Handler {
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
