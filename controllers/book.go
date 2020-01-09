package controllers
 
import (
	"strconv"
	"log"
	"encoding/json"
	"database/sql"
	"net/http"
	"../models"
	"../repository/book"

	"github.com/gorilla/mux"
)

type Controller struct{}

var books []models.Book

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func (c Controller) GetBooks(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter,r *http.Request) {
		var book models.Book
		books = []models.Book{}
		bookRepo := bookRepository.BookRepository{}

		books = bookRepo.GetBooks(db, book, books)

		json.NewEncoder(w).Encode(books)
	}
}

func (c Controller) GetBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter,r *http.Request) {
		var book models.Book
		req := mux.Vars(r)
		id, err := strconv.Atoi(req["id"]) 
		logFatal(err)
		
		bookRepo := bookRepository.BookRepository{}

		book = bookRepo.GetBook(db, book, id)

		json.NewEncoder(w).Encode(book)
	}
}

func (c Controller) AddBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter,r *http.Request) {
		var book models.Book
		var BookId int
		_ = json.NewDecoder(r.Body).Decode(&book)

		bookRepo := bookRepository.BookRepository{}
		
		BookId = bookRepo.AddBook(db, book)

		json.NewEncoder(w).Encode(BookId)
	}
}

func (c Controller) UpdateBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter,r *http.Request) {
		var book models.Book
		_ = json.NewDecoder(r.Body).Decode(&book)

		bookRepo := bookRepository.BookRepository{}
		count := bookRepo.UpdateBook(db, book)

		json.NewEncoder(w).Encode(count)

	}
}

func (c Controller) RemoveBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter,r *http.Request) {
		req := mux.Vars(r)
		id, err := strconv.Atoi(req["id"]) 
		logFatal(err)

		bookRepo := bookRepository.BookRepository{}
		count := bookRepo.RemoveBook(db, id)

		json.NewEncoder(w).Encode(count)
	}
}
