package main

import (
	"log"

	"github.com/gorilla/mux"
	"github.com/harik8/todo-list-service/routehandler"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/todo/", addTodoHandler).Methods("POST")
	r.HandleFunc("/todo/", getTodosHandler).Methods("GET")
	r.HandleFunc("/todo/{TID}", getTodoHandler).Methods("GET")
	r.HandleFunc("/todo/{TID}", updateTodoHandler).Methods("PUT")
	r.HandleFunc("/todo/{TID}", deleteTodoHandler).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", r))
}