package main

import (
	"net/http"
	"todolist/handler"
	"todolist/middleware"
)

func main() {
	runRouter()
}

func runRouter() {
	mutex := http.NewServeMux()

	mutex.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		http.Redirect(writer, request, "/todolist", http.StatusPermanentRedirect)
	})
	mutex.HandleFunc("/todolist", handler.IndexHandler)

	mutex.HandleFunc("/todolist/create", handler.TodolistHandlerCreate)
	mutex.HandleFunc("/todolist/create/post", handler.TodolistHandlerNewTodolist)

	mutex.HandleFunc("/todolist/get/", handler.TodolistHandlerGetTodolist)

	mutex.HandleFunc("/todolist/delete", handler.TodolistHandlerDeleteTodolist)

	errorMiddleware := middleware.ErrorMiddleware{
		Handler: mutex,
	}

	server := http.Server{
		Addr: ":8080",
		Handler: errorMiddleware,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
