package main

import (
	"car_catalog/pkg/models/postgres"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

func main() {
	db, err := postgres.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	mux := http.NewServeMux()
	mux.HandleFunc("/cars", getCars)
	mux.HandleFunc("/cars/delete", deleteCars)

	log.Println("Starting server on :4000")
	err = http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
