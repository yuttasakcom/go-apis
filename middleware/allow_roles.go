package middleware

import (
	"net/http"
)

// AllowRoles Middleware
func AllowRoles(roles ...string) Middleware {
	allow := make(map[string]struct{})

	for _, role := range roles {
		allow[role] = struct{}{}
	}

	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			reqRole := r.Header.Get("Role")
			if _, ok := allow[reqRole]; !ok {
				http.Error(w, "Forbidden", http.StatusForbidden)
				return
			}
			h.ServeHTTP(w, r)
		})
	}
}
