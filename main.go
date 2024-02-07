package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r1 := mux.NewRouter()

	println("Server starting on the port 8080....")

	err := http.ListenAndServe(":8080", r1)
	if err != nil {
		log.Fatal("Error running the server:", err)
	}

}
