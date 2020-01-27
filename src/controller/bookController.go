package controller

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"models"
	"net/http"
)

type Controller struct{}

func (c Controller) GetBooks(db *sql.DB) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json")
		var book models.Book
		books := []models.Book{}
		rows, _ := db.Query("select * from book")
		defer rows.Close()
		//defer handleInternalServerError()
		//panicOnError(err)
		for rows.Next() {
			err := rows.Scan(&book.Id, &book.Name, &book.Author, &book.Year)
			logError(err)
			//	panicOnError(err)
			books = append(books, book)
		}
		json.NewEncoder(writer).Encode(books)
	}
}

func (c Controller) AddBook(db *sql.DB) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json")
		var book models.Book
		json.NewDecoder(request.Body).Decode(&book)
		query := "insert into book (title, author, year) values ($1,$2,$3) returning id"
		row := db.QueryRow(query, book.Name, book.Author, book.Year)
		err := row.Scan(&book.Id)
		logError(err)
		//panicOnError(err)
		json.NewEncoder(writer).Encode(book)
	}
}

func (c Controller) UpdateBook(db *sql.DB) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json")
		var book models.Book
		json.NewDecoder(request.Body).Decode(&book)
		query := "update book set title=$1 , author=$2, year=$3 where id = $4 returning id"
		result, err := db.Exec(query, book.Name, book.Author, book.Year, book.Id)
		logError(err)
		//panicOnError(err)
		rowsUpdated, err := result.RowsAffected()
		logError(err)
		//panicOnError(err)
		json.NewEncoder(writer).Encode(rowsUpdated)
	}
}

func (c Controller) DeleteBook(db *sql.DB) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json")
		params := mux.Vars(request)
		bookId := params["id"]
		query := "delete from book where id = $1"
		result, err := db.Exec(query, bookId)
		logError(err)
		//panicOnError(err)
		rowsAffected, _ := result.RowsAffected()
		json.NewEncoder(writer).Encode(rowsAffected)
	}
}

func (c Controller) GetBookById(db *sql.DB) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json")
		params := mux.Vars(request)
		bookId := params["id"]
		var book models.Book
		row := db.QueryRow("Select * from book where id = $1", bookId)
		err := row.Scan(&book.Id, &book.Name, &book.Author, &book.Year)
		logError(err)
		//panicOnError(err)
		json.NewEncoder(writer).Encode(book)
	}
}

func logError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

/*func handleInternalServerError() {
	if r := recover(); r != nil {
		log.Println("Error recovered", r)
	}
}*/

func errorHandler(writer http.ResponseWriter, request *http.Request, status int) {
	writer.WriteHeader(status)
	if status == http.StatusNotFound {
		error := models.ErrorMessage{"Id not Found"}
		json.NewEncoder(writer).Encode(error)
	}
}

/*
func panicOnError(err error) {
	if err != nil {
		panic(err)
	}
}*/
