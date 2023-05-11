package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/moonidelight/go_course/lab3/internal/delivery"
	"github.com/moonidelight/go_course/lab3/internal/handler"
	"github.com/moonidelight/go_course/lab3/internal/repository"
	"github.com/moonidelight/go_course/lab3/internal/usecase"
	"net/http"
)

func main() {
	repo := repository.NewBookRepository()
	if repo == nil {
		return
	}
	bu := usecase.NewPsqlImp(repo)
	bd := delivery.NewBookDelivery(bu)
	Router := mux.NewRouter().StrictSlash(true)

	handler.Handler(Router, bd)

	fmt.Println("Server is running...")
	err := http.ListenAndServe(":8181", Router)
	if err != nil {
		fmt.Println(err.Error())
	}

}
