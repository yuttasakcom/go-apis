package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/yuttasakcom/go-apis/handlers"
	"github.com/yuttasakcom/go-apis/middleware"
)

// Router http.Handler
func Router() http.Handler {

	// Gorilla mux
	r := mux.NewRouter()

	// Auth handler
	r.HandleFunc("/login", handlers.AuthLogin).Methods("POST")

	// Users handler
	r.Handle("/users", middleware.Chain(
		middleware.Auth("token"),
		// middleware.AllowRoles("admin", "staff"),
	)(http.HandlerFunc(handlers.UserAll))).Methods("GET")
	r.HandleFunc("/users/{id}", handlers.UserID).Methods("GET")
	r.HandleFunc("/users", handlers.UserCreate).Methods("POST")
	r.HandleFunc("/users/{id}", handlers.UserUpdate).Methods("PUT")
	r.HandleFunc("/users/{id}", handlers.UserDelete).Methods("DELETE")

	// HealthCheck handler
	r.HandleFunc("/health", handlers.Health).Methods("GET")

	// Global middleware
	r.Use(middleware.RateLimit) // Todo: change to the Redis
	r.Use(middleware.Logging)

	return r
}
