package routehandler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"

	exec "github.com/harik8/todo-list-service/dbexecutor"
)

// func AddTodoHandler(w http.ResponseWriter, r *http.Request) {
// 	/**
// 		https://golang.org/pkg/encoding/json/#NewDecoder
// 	**/
// 	var td exec.Todo

// 	w.Header().Set("content-type", "application/json")
// 	json.NewDecoder(r.Body).Decode(&td)
// 	exec.AddTodo(td)
// 	json.NewEncoder(w).Encode(td)
// }

func AddTodoHandler(w http.ResponseWriter, r *http.Request) {
	/**
			https://golang.org/pkg/encoding/json/#NewDecoder
	**/
	var td exec.Todo

	w.Header().Set("content-type", "application/json")
	json.NewDecoder(r.Body).Decode(&td)
	exec.AddTodo(td)
	json.NewEncoder(w).Encode(td)
}

func UpdateTodoHandler(w http.ResponseWriter, r *http.Request) {
	var td exec.Todo
	w.Header().Set("content-type", "application/json")
	params := mux.Vars(r)
	tid, _ := primitive.ObjectIDFromHex(params["TID"])
	json.NewDecoder(r.Body).Decode(&td)
	err := exec.UpdateTodo(tid, td)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	json.NewEncoder(w).Encode(td)
}

func DeleteTodoHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	params := mux.Vars(r)
	tid, _ := primitive.ObjectIDFromHex(params["TID"])
	d, err := exec.DeleteTodo(tid)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	json.NewEncoder(w).Encode(d)
}

func GetTodoHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	params := mux.Vars(r)
	tid, _ := primitive.ObjectIDFromHex(params["TID"])
	t, err := exec.GetTodo(tid)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	json.NewEncoder(w).Encode(t)
}

func GetTodosHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	t, err := exec.GetTodos()
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	json.NewEncoder(w).Encode(t)
}