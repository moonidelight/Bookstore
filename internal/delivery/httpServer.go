package delivery

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/moonidelight/go_course/lab3/internal/usecase"
	"io/ioutil"
	"net/http"
	"reflect"
	"strconv"
)

type BookDelivery struct {
	bu *usecase.PsqlImp
}

func NewBookDelivery(bu *usecase.PsqlImp) *BookDelivery {
	return &BookDelivery{bu: bu}
}

func (bd *BookDelivery) Create(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var data map[string]interface{}
	err := json.Unmarshal([]byte(string(reqBody)), &data)
	if err != nil {
		return
	}
	bd.bu.Create(data["title"].(string), data["description"].(string), data["cost"].(float64))
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, "Book was added")
	fmt.Println("Endpoint hit: Add book")
}

func (bd *BookDelivery) GetById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key, _ := strconv.Atoi(vars["Id"])
	title, desc, cost := bd.bu.GetById(key)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "Title: %v; Description: %v; Cost: %v", title, desc, cost)
	fmt.Println("Endpoint hit: GetById")

}

func (bd *BookDelivery) GetAll(w http.ResponseWriter, r *http.Request) {
	books := bd.bu.GetAll()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
	fmt.Println("Endpoint hit: GetAll")
}

func (bd *BookDelivery) UpdateTitleAndDesc(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key, _ := strconv.Atoi(vars["Id"])
	reqBody, _ := ioutil.ReadAll(r.Body)
	var data map[string]interface{}
	err := json.Unmarshal([]byte(string(reqBody)), &data)
	if err != nil {
		return
	}
	title, description := bd.bu.UpdateTitleAndDesc(key, data["title"].(string), data["description"].(string))
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "book with id %v updated to new title %v and description %v", key, title, description)
	fmt.Println("Endpoint hit: UpdateTitleAndDescription")
}

func (bd *BookDelivery) DeleteById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["Id"])
	bd.bu.DeleteById(id)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "Book with id %v was deleted", id)
	fmt.Println("Endpoint hit: DeleteById")
}

func (bd *BookDelivery) SearchByTitle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key, _ := vars["Title"]
	fmt.Println(reflect.TypeOf(key))
	title, desc, cost := bd.bu.SearchByTitle(key)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "Title: %v; Description: %v; Cost: %v", title, desc, cost)
	fmt.Println("Endpoint hit: SearchByTitle")
}

func (bd *BookDelivery) GetSortedBooksAsc(w http.ResponseWriter, r *http.Request) {
	sortedBooks := bd.bu.GetSortedBooksAsc()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(sortedBooks)
	fmt.Println("Endpoint hit: GetSortedBooksAsc")
}

func (bd *BookDelivery) GetSortedBooksDesc(w http.ResponseWriter, r *http.Request) {
	sortedBooks := bd.bu.GetSortedBooksDesc()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(sortedBooks)
	fmt.Println("Endpoint hit: GetSortedBooksDesc")
}
