package middleware

import (
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Authorization header required", http.StatusUnauthorized)
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte("7f1b7e0d3c6a4b2e8f5d9c3b5a8e6f4d7c1b3e5a9f2d4c6b8e1f3a5d7c9b2e4"), nil 
		})

		if err != nil || !token.Valid {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	}
}


func AdminMiddleware(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        tokenString := strings.TrimPrefix(r.Header.Get("Authorization"), "Bearer ")
        token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
            return []byte("your-256-bit-secret"), nil
        })

        claims := token.Claims.(jwt.MapClaims)
        if claims["role"] != "admin" {
            http.Error(w, "Admin access required", http.StatusForbidden)
            return
        }

        next(w, r)
    }
}