package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/matheusosorio/orm-server/pkg/controller"
)

var router *mux.Router

func Start(port string) {
	router = mux.NewRouter()
	setupRoutes()

	http.ListenAndServe(":"+port, router)
}

func setupRoutes() {
	router.HandleFunc("/book", controller.CreateBook).Methods("POST")            // Creates book
	router.HandleFunc("/book", controller.GetBook).Methods("GET")                // returns all books
	router.HandleFunc("/book/{bookId}", controller.GetBookById).Methods("GET")   // returns specific book
	router.HandleFunc("/book/{bookId}", controller.UpdateBook).Methods("PUT")    // updates book
	router.HandleFunc("/book/{bookId}", controller.DeleteBook).Methods("DELETE") // deletes book
}
