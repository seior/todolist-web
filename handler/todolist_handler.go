package handler

import (
	"context"
	"net/http"
	"strconv"
	"todolist/model"
	"todolist/repository"
	"todolist/service"
	"todolist/utils"
	"todolist/view"
)

func TodolistHandlerCreate(w http.ResponseWriter, r *http.Request) {
	err := view.ViewHTML.ExecuteTemplate(w, "create_todolist.gohtml", map[string]interface{}{
		"Title": "Create New Todolist",
	})
	if err != nil {
		panic(err)
	}
}

func TodolistHandlerNewTodolist(w http.ResponseWriter, r *http.Request) {
	todolistRepository := repository.NewTodolistRepository(utils.GetDatabase())

	err := r.ParseForm()
	if err != nil {
		panic(err)
	}

	todolist := model.Todolist{
		Name:   r.FormValue("name"),
		Author: r.FormValue("author"),
	}

	if todolist.Name != "" && todolist.Author != "" {
		ctx := context.Background()
		service.NewTodolistService(todolistRepository).CreateTodolist(ctx, &todolist)
	}

	http.Redirect(w, r, "/", http.StatusPermanentRedirect)
}

func TodolistHandlerGetTodolist(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	todolistRepository := repository.NewTodolistRepository(utils.GetDatabase())

	err := view.ViewHTML.ExecuteTemplate(w, "get_todolist.gohtml", map[string]interface{}{
		"Title":    "List Todolist",
		"Todolist": service.NewTodolistService(todolistRepository).ShowTodolist(ctx),
	})

	if err != nil {
		panic(err)
	}
}

func TodolistHandlerDeleteTodolist(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	todolistRepository := repository.NewTodolistRepository(utils.GetDatabase())

	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		panic(err)
	}

	service.NewTodolistService(todolistRepository).DeleteTodolist(ctx, id)

	http.Redirect(w, r, "/todolist/get", http.StatusTemporaryRedirect)
}
