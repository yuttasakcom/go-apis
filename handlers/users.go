package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/yuttasakcom/go-apis/models"
)

// UserAll handler
func UserAll(w http.ResponseWriter, _ *http.Request) {

}

// UserID handler
func UserID(w http.ResponseWriter, r *http.Request) {

}

// UserCreate handler
func UserCreate(w http.ResponseWriter, r *http.Request) {
	var u models.User
	_ = json.NewDecoder(r.Body).Decode(&u)
	u.Create()
}

// UserUpdate handler
func UserUpdate(w http.ResponseWriter, r *http.Request) {

}

// UserDelete handler
func UserDelete(w http.ResponseWriter, r *http.Request) {

}
