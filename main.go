package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	rh "github.com/harik8/todo-list-service/routehandler"
	conf "github.com/harik8/todo-list-service/config"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/todo/", 		rh.AddTodoHandler).Methods("POST")
	r.HandleFunc("/todo/", 		rh.GetTodosHandler).Methods("GET")
	r.HandleFunc("/todo/{TID}", rh.GetTodoHandler).Methods("GET")
	r.HandleFunc("/todo/{TID}", rh.UpdateTodoHandler).Methods("PUT")
	r.HandleFunc("/todo/{TID}", rh.DeleteTodoHandler).Methods("DELETE")
	log.Println("Starting ToDo Service........")
	err := http.ListenAndServe(conf.TodoServiceIP + ":" + conf.TodoServicePort, r)
	if err != nil {
		panic(err)
	}
}