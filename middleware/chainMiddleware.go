package middleware

import "net/http"

// Middleware func(http.Handler) http.Handler
type Middleware func(http.Handler) http.Handler

// Chain Middleware
func Chain(hs ...Middleware) Middleware {
	return func(h http.Handler) http.Handler {
		for i := len(hs); i > 0; i-- {
			h = hs[i-1](h)
		}
		return h
	}
}
