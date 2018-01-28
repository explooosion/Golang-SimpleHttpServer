package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", helloHandler)

	http.ListenAndServe(":8080", RecoveryHandler(r))
}

func helloHandler(w http.ResponseWriter, req *http.Request) {
	//panic("Die")

	pass := true
	if pass {
		panic("Die")
	}
	io.WriteString(w, "Hello 讀書會")
}

func RecoveryHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				w.WriteHeader(http.StatusInternalServerError)
				io.WriteString(w, fmt.Sprintf("Internal error: %v", r))
			}
		}()

		next.ServeHTTP(w, req)
	})
}
