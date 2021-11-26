package middleware

import (
	"fmt"
	"net/http"
)

type ErrorMiddleware struct {
	Handler http.Handler
}

func (middleware ErrorMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("Error Occured")
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Error : %s", err)
		}
	}()
	middleware.Handler.ServeHTTP(w, r)
}