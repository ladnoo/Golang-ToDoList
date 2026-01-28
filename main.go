package main

import (
	"ToDoList/httpserver"
	"ToDoList/todo"
	"fmt"
)

func main() {
	todoList := todo.NewList()
	httpHandlers := httpserver.NewHTTPHandlers(todoList)
	httpServer := httpserver.NewServer(httpHandlers)

	if err := httpServer.StartServer(); err != nil {
		fmt.Println("Failed to start HTTP server:", err)
	}
}
