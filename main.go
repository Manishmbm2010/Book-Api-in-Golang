package main

import (
	"controller"
	"database/sql"
	"driver"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/subosito/gotenv"
	"log"
	"net/http"
)

var db *sql.DB

func init() {
	gotenv.Load()
}

func main() {
	db = driver.ConnectDb()
	c := controller.Controller{}
	router := mux.NewRouter()
	router.HandleFunc("/", helloServer).Methods("GET")
	router.HandleFunc("/books", c.GetBooks(db)).Methods("GET")
	router.HandleFunc("/book/{id}", c.GetBookById(db)).Methods("GET")
	router.HandleFunc("/book/{id}", c.DeleteBook(db)).Methods("DELETE")
	router.HandleFunc("/book", c.AddBook(db)).Methods("POST")
	router.HandleFunc("/book", c.UpdateBook(db)).Methods("PUT")
	fmt.Println("Server started")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func helloServer(writer http.ResponseWriter, request *http.Request) {
	json.NewEncoder(writer).Encode("Hello from Go World")
}



Hi Kavish , how are u doing ?
matched the horoscope ?
if its okay we can take this further, could you please let me know the outcome
