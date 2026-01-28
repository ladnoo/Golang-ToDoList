package main

import (
	transport2 "ToDoList/internal/transport"
	"ToDoList/todo"
	"fmt"
)

func main() {

	// replace "todoList" with db
	todoList := todo.NewList()
	httpHandlers := transport2.NewHTTPHandlers(todoList)
	httpServer := transport2.NewServer(httpHandlers)

	if err := httpServer.StartServer(); err != nil {
		fmt.Println("Failed to start HTTP server:", err)
	}
}
