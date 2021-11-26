package handler

import (
	"net/http"
	"todolist/view"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	err := view.ViewHTML.ExecuteTemplate(w, "index.gohtml", map[string]interface{}{
		"Title" : "Todolist App",
	})
	if err != nil {
		panic(err)
	}
}
