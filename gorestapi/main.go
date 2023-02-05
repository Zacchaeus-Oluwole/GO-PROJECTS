package main

import (
	"fmt"
	"encoding/json"
	"log"
	"net/http"
	// "math/rand"
	// "strconv"
	"github.com/gorilla/mux"
	"gorestapi/src"
)

//Init books var as a slice Book struct
var books []models.Book

//Get all books
func getBooks(w http.ResponseWriter, r *http.Request){
	fmt.Println("Status-Code 200 Successful...")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func main()  {
	//Init Router
	r := mux.NewRouter()

	//Mock Data - @todo - implement DB

	books = append(books, models.Book{ID: "1", Isbn: "12345", 
	Title: "First Book", Author: &models.Author{Firstname: "Zacchaeus", Lastname: "Oluwole"}})

	books = append(books, models.Book{ID: "2", Isbn: "34567", 
	Title: "Second Book", Author: &models.Author{Firstname: "Olujimi", Lastname: "Oluwole"}})

	// Route Handlers / Endpoints
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	// r.HandleFunc("api/books/{id}", getBook).Methods("GET")
	// r.HandleFunc("api/books", createBook).Methods("POST")
	// r.HandleFunc("api/books/{id}", updateBook).Methods("PUT")
	// r.HandleFunc("api/books/{id}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))
}