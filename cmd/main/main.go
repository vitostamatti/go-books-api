package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/vitostamatti/books-api/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookRoutes(r)
	http.Handle("/", r)

	httpPort := os.Getenv("API_PORT")
	if httpPort == "" {
		httpPort = "8000"
	}

	log.Fatal(http.ListenAndServe(":"+httpPort, r))
}
