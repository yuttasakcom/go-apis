# go-apis

* [Main](#main)
* [Router](#router)
* [Chain Middleware](#chain-middleware)
* [Auth Middleware](#auth-middleware)
* [Users Handler](#users-handler)

## Main

```go
package main

import (
	"log"
	"net/http"

	"github.com/yuttasakcom/go-apis/routes"
)

func main() {
	r := routes.Router()
	log.Println("go-apis running at port:3000")
	log.Fatal(http.ListenAndServe(":3000", r))
}
```

## Router

```go
package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/yuttasakcom/go-apis/handlers"
	"github.com/yuttasakcom/go-apis/middleware"
)

// Router http.Handler
func Router() http.Handler {

	// Gorilla mux
	r := mux.NewRouter()

	// Auth handler
	r.HandleFunc("/login", handlers.AuthLogin).Methods("POST")

	// Users handler
	r.Handle("/users", middleware.Chain(
		middleware.LoggingMiddleware,
		middleware.AuthMiddleware("token"),
		// middleware.AllowRolesMiddleware("admin", "staff"),
	)(http.HandlerFunc(handlers.UserAll))).Methods("GET")
	r.HandleFunc("/users/{id}", handlers.UserID).Methods("GET")
	r.HandleFunc("/users", handlers.UserCreate).Methods("POST")
	r.HandleFunc("/users/{id}", handlers.UserUpdate).Methods("PUT")
	r.HandleFunc("/users/{id}", handlers.UserDelete).Methods("DELETE")

	return r
}
```

## Chain Middleware

```go
package middleware

import "net/http"

// Middleware func(http.Handler) http.Handler
type Middleware func(http.Handler) http.Handler

// Chain Middleware
func Chain(hs ...Middleware) Middleware {
	return func(h http.Handler) http.Handler {
		for i := len(hs); i > 0; i-- {
			h = hs[i-1](h)
		}
		return h
	}
}
```

## Auth Middleware

```go
package middleware

import (
	"net/http"
)

// AuthMiddleware http.Handler
func AuthMiddleware(token string) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			// Todo auth validate

			next.ServeHTTP(w, r)
		})
	}
}
```

## Users Handler

```go
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
func UserAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf8")
	json.NewEncoder(w).Encode(users)
}

// UserID handler
func UserID(w http.ResponseWriter, r *http.Request) {
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

// UserCreate handler
func UserCreate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf8")

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

// UserDelete handler
func UserDelete(w http.ResponseWriter, r *http.Request) {
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
```
