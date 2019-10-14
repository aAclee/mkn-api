package middleware

import "net/http"

// HandlerFunc allows multiple http.HandlerFunc functions to be applied
// as middleware.
func HandlerFunc(h http.HandlerFunc, middleware ...func(http.HandlerFunc) http.HandlerFunc) http.HandlerFunc {
	for _, m := range middleware {
		h = m(h)
	}

	return h
}
