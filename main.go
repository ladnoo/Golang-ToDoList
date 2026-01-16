package main

import (
	"ToDoList/todo"
	"net/http"
	"time"
)

var thingsToDo = make([string]todo.Task, 0)

func CreateTask(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {

	}

}

func main() {
	http.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {})
}
