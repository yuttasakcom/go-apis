package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/yuttasakcom/go-apis/handlers"
)

// Router handler
func Router() http.Handler {
	r := mux.NewRouter()

	users := handlers.UsersHandler{}
	r.HandleFunc("/users", users.All()).Methods("GET")
	r.HandleFunc("/users/{id}", users.GetByID()).Methods("GET")

	return r
}
