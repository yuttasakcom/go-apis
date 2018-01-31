package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func init() {
	fmt.Println("AuthHandler")
}

// AuthHandler handler
type AuthHandler struct{}

// Login handler
func (AuthHandler) Login() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf8")

		var user User
		_ = json.NewDecoder(r.Body).Decode(&user)

		for _, u := range users {
			if u.Email == user.Email {
				err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(user.Password))
				if err == nil {
					json.NewEncoder(w).Encode(u)
				}

			}
		}

	}
}
