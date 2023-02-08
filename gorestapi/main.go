package main

import (
	// "fmt"
	"encoding/json"
	"log"
	"net/http"
	// "math/rand"
	"strconv"
	"github.com/gorilla/mux"
	"gorestapi/src"
	// "reflect"
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
			// fmt.Println(reflect.TypeOf(item.ID))
			return
		}
	}
	json.NewEncoder(w).Encode(&models.Book{})
}

//Create a New Book
func createBook(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	var book models.Book
	newID := 0
	_ = json.NewDecoder(r.Body).Decode(&book)
	for iNum, item := range books {
		if iNum == len(books) - 1 {	
			newID,_ = strconv.Atoi(item.ID)
		}
	}
	book.ID = strconv.Itoa(newID + 1)
	books = append(books, book)
	json.NewEncoder(w).Encode(book)
}

func deleteBook(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Gets params

	for index, item := range books{
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(books)
}

func updateBook( w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Gets params

	for index, item := range books{
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			var book models.Book
			_ = json.NewDecoder(r.Body).Decode(&book)
			book.ID = params["id"]
			books = append(books, book)
			return
		}
	}
	json.NewEncoder(w).Encode(books)
}


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
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))


}

// {
// 	"isbn":"4545454",
// 	"title":"Book Three",
// 	"author":{"firstname":"Harry","lastname":"White"}
// }