package handler

import (
	"context"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
	"todolist/model"
	"todolist/repository"
	"todolist/utils"
)

func TestMain(m *testing.M) {
	ctx := context.Background()
	todolistRepository := repository.NewTodolistRepository(utils.GetDatabase())
	_, err := todolistRepository.DeleteAll(ctx)
	if err != nil {
		panic(err)
	}

	m.Run()

	_, err = todolistRepository.DeleteAll(ctx)
	if err != nil {
		panic(err)
	}
}

func TestTodolistHandlerCreate(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "localhost:8080/todolist/create", nil)
	recorder := httptest.NewRecorder()

	TodolistHandlerCreate(recorder, request)

	body, err := io.ReadAll(recorder.Result().Body)
	if err != nil {
		panic(err)
	}

	assert.True(t, strings.Contains(string(body), "<body>"))
	assert.True(t, strings.Contains(string(body), "</label>"))
	assert.True(t, strings.Contains(string(body), "</form>"))
	assert.True(t, strings.Contains(string(body), "<title>"))
	assert.True(t, strings.Contains(string(body), "</body>"))
}

func TestTodolistHandlerNewTodolist(t *testing.T) {
	body := strings.NewReader("name=Learn Golang&author=Raden Mohamad Rishwan")
	request := httptest.NewRequest(http.MethodPost, "localhost:8080/todolist/create/post", body)
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	recorder := httptest.NewRecorder()

	TodolistHandlerNewTodolist(recorder, request)

	ctx := context.Background()
	todolistRepository := repository.NewTodolistRepository(utils.GetDatabase())
	result, err := todolistRepository.Get(ctx)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, 1, len(result))
	assert.Equal(t, http.StatusPermanentRedirect, recorder.Code)
}

func TestTodolistHandlerGetTodolist(t *testing.T) {
	todolist := model.Todolist{
		Name:   "Learn Golang",
		Author: "Raden Mohamad Rishwan",
	}

	ctx := context.Background()
	todolistRepository := repository.NewTodolistRepository(utils.GetDatabase())
	todolistRepository.Create(ctx, &todolist)
	todolistRepository.Create(ctx, &todolist)

	request := httptest.NewRequest(http.MethodGet, "localhost:8080/todolist/get", nil)
	recorder := httptest.NewRecorder()

	TodolistHandlerGetTodolist(recorder, request)

	body, err := io.ReadAll(recorder.Result().Body)
	if err != nil {
		panic(err)
	}

	assert.True(t, strings.Contains(string(body), "<body>"))
	assert.True(t, strings.Contains(string(body), "<ul>"))
	assert.True(t, strings.Contains(string(body), "<li>"))
	assert.True(t, strings.Contains(string(body), "<p>"))
	assert.True(t, strings.Contains(string(body), "delete"))
	assert.True(t, strings.Contains(string(body), "</body>"))
}

func TestTodolistHandlerDeleteTodolist(t *testing.T) {
	todolist := model.Todolist{
		Name:   "Learn Golang",
		Author: "Raden Mohamad Rishwan",
	}

	ctx := context.Background()
	todolistRepository := repository.NewTodolistRepository(utils.GetDatabase())
	id, err := todolistRepository.Create(ctx, &todolist)
	if err != nil {
		panic(err)
	}

	request := httptest.NewRequest(http.MethodGet, "localhost:8080/todolist/delete?id="+strconv.Itoa(id), nil)
	recorder := httptest.NewRecorder()

	TodolistHandlerDeleteTodolist(recorder, request)

	get, err := todolistRepository.GetByIndex(ctx, uint32(id))
	if err != nil {
		panic(err)
	}

	assert.Equal(t, 0, int(get.Id))
	assert.Equal(t, "", get.Name)
	assert.Equal(t, "", get.Author)
}
