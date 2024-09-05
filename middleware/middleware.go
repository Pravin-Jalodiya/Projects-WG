package middleware

import (
	"context"
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"projects/utils/errs"
	"projects/utils/logger"
)

var secretKey = []byte("Xwdwq0a1da3sqe20awas0e-qwe0dq0wd032-qd0da0sdas02-ascas0cas0f")

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
			err := errs.NewUnauthorizedError("Authorization token required")
			err.ToJSON(w)
			return
		}

		token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
			return secretKey, nil
		})

		if err != nil {
			logger.Logger.Errorw("Token parsing error",
				"error", err,
				"time", time.Now())
			errMsg := "Invalid token"
			var ve *jwt.ValidationError
			if errors.As(err, &ve) {
				if ve.Errors&jwt.ValidationErrorExpired != 0 {
					errMsg = "Token has expired"
				}
			}
			err := errs.NewUnauthorizedError(errMsg)
			err.ToJSON(w)
			return
		}

		claims, ok := token.Claims.(*Claims)
		if !ok || !token.Valid {
			err := errs.NewUnauthorizedError("Invalid token claims")
			err.ToJSON(w)
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
