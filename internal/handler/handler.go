package handler

import (
	"github.com/gorilla/mux"
	"github.com/moonidelight/go_course/lab3/internal/delivery"
)

func Handler(Router *mux.Router, bd *delivery.BookDelivery) {
	Router.HandleFunc("/NewBook", bd.Create).Methods("POST")
	Router.HandleFunc("/books/{Id}", bd.GetById).Methods("GET")
	Router.HandleFunc("/books", bd.GetAll).Methods("GET")
	Router.HandleFunc("/books/{Id}", bd.UpdateTitleAndDesc).Methods("PUT")
	Router.HandleFunc("/books/{Id}", bd.DeleteById).Methods("DELETE")
	Router.HandleFunc("/books/title/{Title}", bd.SearchByTitle).Methods("GET")
	Router.HandleFunc("/sorted_books_asc", bd.GetSortedBooksAsc).Methods("GET")
	Router.HandleFunc("/sorted_books_desc", bd.GetSortedBooksDesc).Methods("GET")
}
