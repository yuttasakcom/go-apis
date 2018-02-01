package middleware

import (
	"log"
	"net/http"
)

// LoggingMiddleware http.Handler
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.RequestURI, "LoggingMiddleware")
		next.ServeHTTP(w, r)
	})
}
