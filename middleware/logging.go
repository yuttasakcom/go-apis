package middleware

import (
	"log"
	"net/http"
)

// Logging http.Handler
func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.RequestURI, "LoggingMiddleware")
		next.ServeHTTP(w, r)
	})
}
