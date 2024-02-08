package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	handlers "github.com/gowdaganesh005/stocksapi/middleware"
)

func main() {
	r1 := mux.NewRouter()

	_, err := handlers.Connect()
	if err != nil {
		log.Println("Error connecting the database", err)
	}

	r1.HandleFunc("/stocks", handlers.CreateStock).Methods("POST")
	r1.HandleFunc("/stocks/id", handlers.GetStock).Methods("GET")

	println("Server starting on the port 8080....")

	err = http.ListenAndServe(":8080", r1)
	if err != nil {
		log.Fatal("Error running the server:", err)
	}

}
