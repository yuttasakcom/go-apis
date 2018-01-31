package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/yuttasakcom/go-apis/handlers"
)

// Router handler
func Router() http.Handler {
	r := mux.NewRouter()

	r.HandleFunc("/users", handlers.UsersHandler())

	return r
}
