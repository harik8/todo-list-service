package routehandler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/harik8/todo-list-service/dbexecutor"
)

func addTodoHandler(w http.ResponseWriter, r *http.Request) {
	/**
		https://golang.org/pkg/encoding/json/#NewDecoder
	**/
	var td todo

	w.Header().Set("content-type", "application/json")
	json.NewDecoder(r.Body).Decode(&td)
	result := addTodo(td)
	json.NewEncoder(w).Encode(td)
	log.Println(result.InsertedID)
}

func updateTodoHandler(w http.ResponseWriter, r *http.Request) {
	var td todo
	w.Header().Set("content-type", "application/json")
	params := mux.Vars(r)
	tid, _ := primitive.ObjectIDFromHex(params["TID"])
	json.NewDecoder(r.Body).Decode(&td)
	err := updateTodo(tid, td)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	json.NewEncoder(w).Encode(td)
}

func deleteTodoHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	params := mux.Vars(r)
	tid, _ := primitive.ObjectIDFromHex(params["TID"])
	d, err := deleteTodo(tid)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	json.NewEncoder(w).Encode(d)
}

func getTodoHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	params := mux.Vars(r)
	tid, _ := primitive.ObjectIDFromHex(params["TID"])
	t, err := getTodo(tid)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	json.NewEncoder(w).Encode(t)
}

func getTodosHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	t, err := getTodos()
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	json.NewEncoder(w).Encode(t)
}