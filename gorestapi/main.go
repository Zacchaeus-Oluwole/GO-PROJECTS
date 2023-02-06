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
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

// Get single book
func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Gets params
	// Loop through books and find one with the id from the params
	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&models.Book{})
}

// //Create a New Book
// func createBook(w http.ResponseWriter, r *http.Request){
// 	w.Header().Set("Content-Type", "application/json")
// 	var book Book
// 	_ = json.NewDecoder(r.Body).Decode(&book)
// 	for _, item := range books {
// 		if item.ID == params["id"] {	
// 			json.NewEncoder(w).Encode(item)
// 			return
// 		}
// 	}
// 	book.ID = strconv.Itoa(books.len() + 1)
// }



func main()  {
	//Init Router
	r := mux.NewRouter()

	//Mock Data - @todo - implement DB
	//Init books var as a slice Book struct
	// var books []models.Book

	books = append(books, models.Book{ID: "1", Isbn: "12345", 
	Title: "First Book", Author: &models.Author{Firstname: "Zacchaeus", Lastname: "Oluwole"}})

	books = append(books, models.Book{ID: "2", Isbn: "34567", 
	Title: "Second Book", Author: &models.Author{Firstname: "Olujimi", Lastname: "Oluwole"}})

	// Route Handlers / Endpoints
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	// r.HandleFunc("api/books", createBook).Methods("POST")
	// r.HandleFunc("api/books/{id}", updateBook).Methods("PUT")
	// r.HandleFunc("api/books/{id}", deleteBook).Methods("DELETE")
	fmt.Println(len(books)+1)

	log.Fatal(http.ListenAndServe(":8000", r))


}