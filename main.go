package main

import (
	"net/http"

	"github.com/gorilla/mux"
	route "github.com/harik8/todo-list-service/routehandler"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/todo/", 		route.AddTodoHandler).Methods("POST")
	r.HandleFunc("/todo/", 		route.GetTodosHandler).Methods("GET")
	r.HandleFunc("/todo/{TID}", route.GetTodoHandler).Methods("GET")
	r.HandleFunc("/todo/{TID}", route.UpdateTodoHandler).Methods("PUT")
	r.HandleFunc("/todo/{TID}", route.DeleteTodoHandler).Methods("DELETE")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		panic(err)
	}
}