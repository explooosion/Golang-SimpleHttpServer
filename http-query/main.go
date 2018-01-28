package main // package main

import (
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

func articleHandler(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	articleName := vars["name"]
	page := req.URL.Query().Get("p")
	io.WriteString(w, req.Header.Get("User-Agent"))
	io.WriteString(w, "Articles:")
	io.WriteString(w, articleName)
	io.WriteString(w, ", where page: ")
	io.WriteString(w, page)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/articles/{name}", articleHandler)
	http.ListenAndServe(":8080", r)
}
