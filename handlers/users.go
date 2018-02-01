package handlers

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
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

// UserAll handler
func UserAll(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(users)
}

// UserID handler
func UserID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	params := mux.Vars(r)

	for _, user := range users {
		if user.ID == params["id"] {
			json.NewEncoder(w).Encode(user)
			return
		}
	}

	json.NewEncoder(w).Encode(&User{})
}

// UserCreate handler
func UserCreate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	var user User
	_ = json.NewDecoder(r.Body).Decode(&user)

	user.ID = strconv.Itoa(rand.Intn(10000000))

	hpwd, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		return
	}

	user.Password = string(hpwd)
	user.CreatedAt = time.Now()
	fmt.Println(user)
	users = append(users, user)

	json.NewEncoder(w).Encode(users)
}

// UserUpdate handler
func UserUpdate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

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

// UserDelete handler
func UserDelete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	params := mux.Vars(r)

	for i, user := range users {
		if user.ID == params["id"] {
			users = append(users[:i], users[i+1:]...)
			break
		}
	}

	json.NewEncoder(w).Encode(users)
}