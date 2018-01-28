package main // package main

import (
	"io"
	"net/http"
)

func helloHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello wolrd")
}

func hello2Handler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello 讀書會")
}

func main() {
	r := http.NewServeMux()
	r.HandleFunc("/", helloHandler)
	r.HandleFunc("/hello2", hello2Handler)
	r.HandleFunc("/hello3", hello2Handler)
	http.ListenAndServe(":8080", r)
}
