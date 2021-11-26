package service

import (
	"context"
	"todolist/model"
)

type TodolistService interface {
	ShowTodolist(ctx context.Context) *[]model.Todolist
	CreateTodolist(ctx context.Context, todolist *model.Todolist)
	DeleteTodolist(ctx context.Context, id int)
}
