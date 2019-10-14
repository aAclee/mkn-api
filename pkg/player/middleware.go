package player

import (
	"net/http"

	"github.com/aaclee/mkn-api/pkg/encode"
	mknjwt "github.com/aaclee/mkn-api/pkg/jwt"
	"github.com/dgrijalva/jwt-go"
)

// MiddlewareAdmin is the middleware for verifying request admin status
func MiddlewareAdmin(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		claims, ok := ctx.Value(mknjwt.ContextKey("claims")).(jwt.MapClaims)

		if !ok {
			encode.ErrorJSON(w, http.StatusUnauthorized, "admin status required for access")
			return
		}

		admin, ok := claims["adn"].(bool)
		if !ok || !admin {
			encode.ErrorJSON(w, http.StatusUnauthorized, "admin status required for access")
			return
		}

		h.ServeHTTP(w, r)
	}
}
