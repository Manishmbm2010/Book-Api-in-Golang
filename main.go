package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

type Book struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Author string `json:"author"`
	Year   string `json:"year"`
}

type ErrorMessage struct {
	Message string `json:"message"`
}

func initializeBooks() {
	books = []Book{
		Book{1, "Go Basics", "Manish Jain", "2019"},
		Book{2, "Go Intermediate", "Anurag Jain", "2018"},
		Book{3, "Go Advanced", "Pramod Jain", "2019"},
		Book{4, "Go Concurrency", "Nupur Jain", "2020"},
	}
}

var books []Book

func main() {
	initializeBooks()
	router := mux.NewRouter()
	router.HandleFunc("/", helloServer).Methods("GET")
	router.HandleFunc("/books", getBooks).Methods("GET")
	router.HandleFunc("/book/{id}", getBookById).Methods("GET")
	router.HandleFunc("/book/{id}", deleteBookById).Methods("DELETE")
	router.HandleFunc("/book", createBook).Methods("POST")
	router.HandleFunc("/book", updateBook).Methods("PUT")
	fmt.Println("Server started")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func updateBook(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	var book Book
	json.NewDecoder(request.Body).Decode(&book)
	for i, b := range books {
		if b.Id == book.Id {
			books[i] = book
			json.NewEncoder(writer).Encode(books[i])
			return
		}
	}
	errorHandler(writer, request, http.StatusNotFound)
}

func createBook(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	var book Book
	json.NewDecoder(request.Body).Decode(&book)
	books = append(books, book)
	json.NewEncoder(writer).Encode(book)
}

func deleteBookById(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	bookId, _ := strconv.Atoi(params["id"])
	for i, book := range books {
		if book.Id == bookId {
			books = append(books[:i], books[i+1:]...)
			writer.WriteHeader(http.StatusNoContent)
			return
		}
	}
	errorHandler(writer, request, http.StatusNotFound)

}

func getBookById(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	bookId, _ := strconv.Atoi(params["id"])

	for _, book := range books {
		if book.Id == bookId {
			json.NewEncoder(writer).Encode(book)
			return
		}
	}
	errorHandler(writer, request, http.StatusNotFound)
}

func helloServer(writer http.ResponseWriter, request *http.Request) {
	json.NewEncoder(writer).Encode("Hello from Go World")
}

func getBooks(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(books)
}

func errorHandler(writer http.ResponseWriter, request *http.Request, status int) {
	writer.WriteHeader(status)
	if status == http.StatusNotFound {
		error := ErrorMessage{"Id not Found"}
		json.NewEncoder(writer).Encode(error)
	}
}
