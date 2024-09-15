package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

func AuthMiddlware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Authorization header missing", http.StatusUnauthorized)
			return
		}
		tokenString := strings.Split(authHeader, "Bearer ")[1]
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, err) {
			return "kalai1623", nil
		})
		if err != nil {
			return
		}
		if !token.Valid {
			http.Error(w, "iNVALID TOKEN SEND")
			return
		}
		ctx := context.WithValue(r.Context(), "email", token.Claims.Valid())
		next.ServeHTTP(w, r.WithContext(ctx))
	})

}
