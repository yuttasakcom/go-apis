package middleware

import (
	"log"
	"net/http"
)

// AuthMiddleware http.Handler
func AuthMiddleware(token string) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Println("token:", token)
			next.ServeHTTP(w, r)
		})
	}
}
