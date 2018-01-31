package routes

import (
	"net/http"

	"github.com/yuttasakcom/go-apis/middleware"

	"github.com/gorilla/mux"
	"github.com/yuttasakcom/go-apis/handlers"
)

// Router handler
func Router() http.Handler {
	r := mux.NewRouter()

	auth := handlers.AuthHandler{}
	r.HandleFunc("/login", auth.Login()).Methods("POST")

	users := handlers.UsersHandler{}
	r.Handle("/users", middleware.LoggingMiddleware(http.HandlerFunc(users.All()))).Methods("GET")
	// r.HandleFunc("/users", users.All()).Methods("GET")
	r.HandleFunc("/users/{id}", users.GetByID()).Methods("GET")
	r.HandleFunc("/users", users.Create()).Methods("POST")
	r.HandleFunc("/users/{id}", users.Update()).Methods("PUT")
	r.HandleFunc("/users/{id}", users.Delete()).Methods("DELETE")

	return r
}
