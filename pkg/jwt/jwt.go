package jwt

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/aaclee/mkn-api/pkg/encode"
	"github.com/dgrijalva/jwt-go"
)

// ContextKey is the type for the JWT middleware context
type ContextKey string

const (
	// ClaimsKey is the constant for the context key cliams
	ClaimsKey = ContextKey("claims")
)

// Verify token validitiy and within expiration time
func Verify(encodedToken string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte("munchkin-secret"), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}

// MiddlewareVerify is the middleware for verifying token validity
func MiddlewareVerify(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)

		s := strings.SplitN(r.Header.Get("Authorization"), " ", 2)
		if len(s) != 2 {
			encode.ErrorJSON(w, http.StatusUnauthorized, "Incorrect token format \"Bearer <token>\"")
			return
		}

		tokenString := s[1]
		claims, err := Verify(tokenString)
		if err != nil {
			encode.ErrorJSON(w, http.StatusUnauthorized, err.Error())
			return
		}

		ctx := context.WithValue(context.Background(), ClaimsKey, claims)
		r = r.WithContext(ctx)

		h.ServeHTTP(w, r)
	}
}

// ParseRequest retrieves the claims from the request
func ParseRequest(r *http.Request) (jwt.MapClaims, error) {
	claims, ok := r.Context().Value(ClaimsKey).(jwt.MapClaims)
	if !ok {
		return nil, errors.New("Error parsing claims")
	}

	return claims, nil
}
