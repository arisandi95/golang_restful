package main

import (
	"strconv"
	// "reflect" => buat cek tipe data
	"encoding/json"
	// "fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Book struct {
	ID int `json:id`
	Title string `json:title`
	Author string `json:author`
	Year string `json:year`
}

var books []Book

func main() {
	books = append(books,
		Book{
			ID: 1,
			Title: "The Lean Startup",
			Author: "Eric Ries",
			Year: "2019",
		},
		Book{
			ID: 2,
			Title: "Awali dengan Bismillah, akhiri dengan Alhamdulillah",
			Author: "Muhammad Arisandi",
			Year: "2020",
		},
		Book{
			ID: 3,
			Title: "Buku Nikah",
			Author: "Sands",
			Year: "2020",
		},
	)

	router := mux.NewRouter()

	router.HandleFunc("/books", getBooks).Methods("GET")
	router.HandleFunc("/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/books", addBook).Methods("POST")
	router.HandleFunc("/books", updateBook).Methods("PUT")
	router.HandleFunc("/books/{id}", removeBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}

func getBooks(w http.ResponseWriter,r *http.Request) {
	log.Println("Gets All Books")
	json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter,r *http.Request) {
	log.Println("Gets One Books")
	req := mux.Vars(r)
	
	idBook, _ := strconv.Atoi(req["id"])

	//cek tipe data
	// log.Println(reflect.TypeOf(i), err, req)

	for _, book := range books {
		if book.ID == idBook {
			json.NewEncoder(w).Encode(&book)
		}
	}


}

func addBook(w http.ResponseWriter,r *http.Request) {
	log.Println("Add book")
	
	var newBook Book
	_ = json.NewDecoder(r.Body).Decode(&newBook)
	books = append(books, newBook)

	json.NewEncoder(w).Encode(books)
}

func updateBook(w http.ResponseWriter,r *http.Request) {
	log.Println("Updates a book")
	var book Book
	json.NewDecoder(r.Body).Decode(&book)

	for i, item := range books {
		if item.ID == book.ID {
			books[i] = book
		}
	}

	json.NewEncoder(w).Encode(books)
}

func removeBook(w http.ResponseWriter,r *http.Request) {
	log.Println("Remove a book")
	req := mux.Vars(r)
	idBook, _ := strconv.Atoi(req["id"])

	for i, item := range books {
		if item.ID == idBook {
			books = append(books[:i], books[i+1:]...)
		}
	}

	json.NewEncoder(w).Encode(books)
}