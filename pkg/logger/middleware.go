package logger

import (
	"net/http"
)

// Middleware intercepts the http pipeline and logs the current request
func Middleware(next http.Handler) http.Handler {
	// Setting up logger
	log := CreateLogger()

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Infof("%s: %s", r.Method, r.RequestURI)

		next.ServeHTTP(w, r)
	})
}
