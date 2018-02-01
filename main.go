package main

import (
	"log"
	"net/http"
	"os"

	"github.com/yuttasakcom/go-apis/routes"
)

func main() {
	r := routes.Router()

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	log.Println("go-apis running at port:" + port)
	log.Println("Press CTRL-C to stop")
	log.Fatal(http.ListenAndServe(":"+port, r))
}
