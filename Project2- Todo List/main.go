package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Todo struct {
	Id      int    `josn:"id"`
	Message string `json:"message"`
}

var todo_slice []Todo

func main() {
	r := mux.NewRouter()
	// initially data
	t1 := Todo{1, "taking bath"}
	t2 := Todo{2, "do brush"}
	fmt.Println(t1)
	todo_slice = append(todo_slice, t1)
	todo_slice = append(todo_slice, t2)
	fmt.Println("Starting server at port 3000: ")
	//routing functions
	r.HandleFunc("/", getAllTodos).Methods("GET")
	r.HandleFunc("/", createTodo).Methods("POST")
	r.HandleFunc("/{id}", updateTodo).Methods("PUT")
	r.HandleFunc("/{id}", deleteTodo).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":3000", r))

}

func getAllTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todo_slice)
}

func createTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application-json")
	var todo Todo
	_ = json.NewDecoder(r.Body).Decode(&todo)
	todo.Id = rand.Intn(100000)
	todo_slice = append(todo_slice, todo)
	json.NewEncoder(w).Encode(todo_slice)
}

func updateTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application-json")
	params := mux.Vars(r)
	temp := false
	for idx, val := range todo_slice {
		s := strconv.Itoa(val.Id)
		if s == params["id"] {
			temp = true
			todo_slice = append(todo_slice[:idx], todo_slice[idx+1:]...)
			break
		}
	}
	if temp == false {
		fmt.Fprintf(w, "Sorry! Particular id not found")
	} else {
		var todo Todo
		_ = json.NewDecoder(r.Body).Decode(&todo)
		todo.Id, _ = strconv.Atoi(params["id"])
		todo_slice = append(todo_slice, todo)
		json.NewEncoder(w).Encode(todo_slice)
	}

}

func deleteTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application-json")
	parmas := mux.Vars(r)
	temp := false
	for idx, val := range todo_slice {
		s := strconv.Itoa(val.Id)
		if s == parmas["id"] {
			temp = true
			todo_slice = append(todo_slice[:idx], todo_slice[idx+1:]...)
			break
		}
	}
	if temp == false {
		fmt.Fprintf(w, "Sorry! ID not found")
	}
	json.NewEncoder(w).Encode(todo_slice)
}
