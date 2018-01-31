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
	UpdatedAt time.Time `json:"updated_at"`
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

		json.NewEncoder(w).Encode(&User{})
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

// Update handler
func (UsersHandler) Update() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf8")

		params := mux.Vars(r)

		for i, user := range users {
			if user.ID == params["id"] {
				users = append(users[:i], users[i+1:]...)
				var user User
				_ = json.NewDecoder(r.Body).Decode(&user)
				user.ID = params["id"]
				user.UpdatedAt = time.Now()
				users = append(users, user)
				json.NewEncoder(w).Encode(user)
				return
			}
		}
	}
}

// Delete handler
func (UsersHandler) Delete() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf8")

		params := mux.Vars(r)

		for i, user := range users {
			if user.ID == params["id"] {
				users = append(users[:i], users[i+1:]...)
				break
			}
		}

		json.NewEncoder(w).Encode(users)
	}
}
