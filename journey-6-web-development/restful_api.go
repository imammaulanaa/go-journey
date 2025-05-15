package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Book struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

var books = []Book{
	{ID: "1", Title: "Go Journey"},
	{ID: "2", Title: "Belajar Go"},
}

func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func main() {
	http.HandleFunc("/books", getBooks)
	fmt.Println("Server berjalan di http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
