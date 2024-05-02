package main

import "net/http"

func getCars(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/cars" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Hello!"))
}

func deleteCars(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Delete!"))
}

func addCars(w http.ResponseWriter, r *http.Request) {

}

func updateCars(w http.ResponseWriter, r *http.Request) {

}
