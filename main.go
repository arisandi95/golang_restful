package main

import (
	"log"
	"net/http"
	"database/sql"
	"./models"
	"./driver"
	"./controllers"

	"github.com/gorilla/mux"
	"github.com/subosito/gotenv"
)

var books []models.Book
var db *sql.DB // Database connection pool.

func init() {
	gotenv.Load()
}

func main() {
	db = driver.ConnectDB()
	router := mux.NewRouter()

	controller := controllers.Controller{}

	router.HandleFunc("/books", controller.GetBooks(db)).Methods("GET")
	router.HandleFunc("/books/{id}", controller.GetBook(db)).Methods("GET")
	router.HandleFunc("/books", controller.AddBook(db)).Methods("POST")
	router.HandleFunc("/books", controller.UpdateBook(db)).Methods("PUT")
	router.HandleFunc("/books/{id}", controller.RemoveBook(db)).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}

