package handler

import (
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestIndexHandlerSuccess(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "localhost:8080/", nil)
	recorder := httptest.NewRecorder()

	IndexHandler(recorder, request)

	body, err := io.ReadAll(recorder.Result().Body)
	if err != nil {
		panic(err)
	}

	assert.True(t, strings.Contains(string(body), "<h5>"))
	assert.True(t, strings.Contains(string(body), "</a>"))
	assert.True(t, strings.Contains(string(body), "<body>"))
	assert.True(t, strings.Contains(string(body), "<title>"))
}
