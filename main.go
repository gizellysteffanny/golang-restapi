package main

import(
	"encoding/json"
	"log"
	"net/http"
	// "math/rand"
	// "strconv"
	"github.com/gorilla/mux"
)

// Book Struct (Model)
type Book struct {
	ID			string `json:"id"`
	Isbn		string `json:"isbn"`
	Title		string `json:"title"`
	Author	*Author `json:"author"`
}

// Author Struct
type Author struct {
	FirstName	string `json:"firstname"`
	LastName		string `json:"lastname"`
}

// Init books var as a slice Book struct
var books []Book

// Get All Books
func getBooks(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

// Get Book
func getBook(w http.ResponseWriter, r *http.Request)  {
	
}

// Create a new Book
func createBook(w http.ResponseWriter, r *http.Request)  {
	
}

// Update the Book
func updateBook(w http.ResponseWriter, r *http.Request)  {
	
}

// Delete the Book
func deleteBook(w http.ResponseWriter, r *http.Request)  {
	
}


func main()  {
	// Init Router
	r := mux.NewRouter()

	// Mock Data - @todo - implement DB
	books = append(books, Book{ID: "1", Isbn: "44874", Title: "Book One", Author: &Author{FirstName: "Joe", LastName: "Doe"}})
	books = append(books, Book{ID: "2", Isbn: "55309", Title: "Book Two", Author: &Author{FirstName: "Ken", LastName: "Meachum"}})

	// Route Handlers / Endpoints
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PATCH")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":3000", r))
}
