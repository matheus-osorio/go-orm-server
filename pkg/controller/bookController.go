package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/imdario/mergo"
	"github.com/matheusosorio/orm-server/pkg/models"
	"github.com/matheusosorio/orm-server/pkg/utils"
)

func CreateBook(writer http.ResponseWriter, request *http.Request) {
	book := models.Book{}
	utils.ParseBody(request, &book)
	book.CreateBook()
	writer.WriteHeader(http.StatusOK)
	response, _ := json.Marshal(book)
	writer.Write(response)
}

func GetBook(writer http.ResponseWriter, request *http.Request) {
	writer = utils.SetDefaultHeaders(writer)
	books := models.GetAllBooks()
	response, err := json.Marshal(books)

	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("Internal Server Error"))
		return
	}
	writer = utils.SetDefaultHeaders(writer)
	writer.WriteHeader(http.StatusOK)
	writer.Write(response)
}

func GetBookById(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	bookId := vars["bookId"]
	id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		writer.Write([]byte("Could not convert " + bookId + "to int"))
		return
	}
	writer = utils.SetDefaultHeaders(writer)
	book, _ := models.GetBookById(id)
	response, _ := json.Marshal(book)
	writer.Write(response)
}

func UpdateBook(writer http.ResponseWriter, request *http.Request) {
	bookUpdate := models.Book{}
	utils.ParseBody(request, &bookUpdate)
	vars := mux.Vars(request)
	bookId := vars["bookId"]

	id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		writer.Write([]byte("Could not convert " + bookId + "to int"))
		return
	}

	book, db := models.GetBookById(id)

	mergo.Merge(&book, bookUpdate, mergo.WithOverride)

	db.Save(&book)

	response, _ := json.Marshal(book)

	writer = utils.SetDefaultHeaders(writer)
	writer.WriteHeader(http.StatusOK)
	writer.Write(response)
}

func DeleteBook(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	bookId := vars["bookId"]
	id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		writer.Write([]byte("Could not convert " + bookId + "to int"))
		return
	}
	writer = utils.SetDefaultHeaders(writer)
	book := models.DeleteBookById(id)
	response, _ := json.Marshal(book)
	writer.Write(response)
}
