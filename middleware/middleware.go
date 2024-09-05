package middleware

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"strings"
)

var secretKey = []byte("your_secret_key") // Change this to a secure key

// Claims JWT claims structure
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// AuthMiddleware to protect routes
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := extractToken(r)
		if tokenString == "" {
			http.Error(w, "Authorization token required", http.StatusUnauthorized)
			return
		}

		token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
			return secretKey, nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
			return
		}

		claims, ok := token.Claims.(*Claims)
		if !ok {
			http.Error(w, "Invalid token claims", http.StatusUnauthorized)
			return
		}

		// Store claims in request context for use in handlers
		r = r.WithContext(context.WithValue(r.Context(), "username", claims.Username))

		next.ServeHTTP(w, r)
	})
}

// Extract the token from the Authorization header
func extractToken(r *http.Request) string {
	authHeader := r.Header.Get("Authorization")
	if authHeader != "" && strings.HasPrefix(authHeader, "Bearer ") {
		return strings.TrimPrefix(authHeader, "Bearer ")
	}
	return ""
}
