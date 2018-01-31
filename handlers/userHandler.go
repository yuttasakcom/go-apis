package handlers

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// User struct
type User struct {
	ID        string    `json:"id"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
}

var users []User

// UsersHandler handler
type UsersHandler struct{}

// All handler
func (UsersHandler) All() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf8")
		json.NewEncoder(w).Encode(users)
	}
}

// GetByID handler
func (UsersHandler) GetByID() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf8")

		params := mux.Vars(r)

		for _, user := range users {
			if user.ID == params["id"] {
				json.NewEncoder(w).Encode(user)
				return
			}
		}

		json.NewEncoder(w).Encode(users)
	}
}

// Create handler
func (UsersHandler) Create() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf8")

		var user User
		_ = json.NewDecoder(r.Body).Decode(&user)
		user.ID = strconv.Itoa(rand.Intn(10000000))
		user.CreatedAt = time.Now()
		users = append(users, user)

		json.NewEncoder(w).Encode(users)
	}
}
