package handlers

import (
	"encoding/json"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

// AuthLogin handler
func AuthLogin(w http.ResponseWriter, r *http.Request) {
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
