package handlers

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/yuttasakcom/go-apis/response"
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
	if len(users) == 0 {
		response.JSON(w, make([]string, 0), http.StatusOK)
		return
	}

	response.JSON(w, users, http.StatusOK)
}

// UserID handler
func UserID(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	for _, user := range users {
		if user.ID == params["id"] {
			response.JSON(w, user, http.StatusOK)
			return
		}
	}

	response.Error(w, http.StatusNotFound)
}

// UserCreate handler
func UserCreate(w http.ResponseWriter, r *http.Request) {
	var user User
	_ = json.NewDecoder(r.Body).Decode(&user)

	user.ID = strconv.Itoa(rand.Intn(10000000))

	hpwd, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		return
	}

	user.Password = string(hpwd)
	user.CreatedAt = time.Now()
	users = append(users, user)

	for _, u := range users {
		if u.ID == user.ID {
			response.JSON(w, user, http.StatusOK)
			return
		}
	}

	response.JSON(w, map[string]string{"Error": "User created fail"}, http.StatusInternalServerError)
}

// UserUpdate handler
func UserUpdate(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	for i, user := range users {
		if user.ID == params["id"] {
			users = append(users[:i], users[i+1:]...)
			var user User
			_ = json.NewDecoder(r.Body).Decode(&user)
			user.ID = params["id"]
			user.UpdatedAt = time.Now()
			users = append(users, user)
			response.JSON(w, user, http.StatusOK)
			return
		}
	}
}

// UserDelete handler
func UserDelete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	for i, user := range users {
		if user.ID == params["id"] {
			users = append(users[:i], users[i+1:]...)
			break
		}
	}

	response.JSON(w, map[string]string{"message": "success"}, http.StatusOK)
}
