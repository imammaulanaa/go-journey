package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Selamat datang di Go Journey")
}

func bookHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprintf(w, "Buku dengan ID: %s\n", vars["id"])
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", homeHandler)
	r.HandleFunc("/book/{id}", bookHandler)

	fmt.Println("Server berjalan di http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
