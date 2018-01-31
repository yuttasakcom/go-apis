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
