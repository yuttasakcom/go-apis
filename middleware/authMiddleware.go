package middleware

import (
	"fmt"
	"net/http"
)

// AuthMiddleware http.Handler
func AuthMiddleware(token string) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			// Todo auth validate
			fmt.Println("AuthMiddleware")

			next.ServeHTTP(w, r)
		})
	}
}
