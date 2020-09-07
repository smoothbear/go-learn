package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type Book struct {
	title       string
	content     string
	description string
}

var books []Book

func createBook(w http.ResponseWriter, r *http.Request) {
	var book Book
	json.NewDecoder(r.Body).Decode(&book)
	books = append(books, book)
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(book)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/createBooks", createBook).Methods("POST")

	http.ListenAndServe(":8000", router)
}
