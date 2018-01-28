package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/",
		http.HandlerFunc(jsonHandler),
	).Methods(http.MethodPost)

	http.ListenAndServe(":8080", r)
}

func jsonHandler(w http.ResponseWriter, req *http.Request) {

	defer req.Body.Close()
	b, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}

	v := struct {
		Name string
		Age  int
	}{}

	if err := json.Unmarshal(b, &v); err != nil {
		panic(err)
	}

	x := struct {
		Type string
		Name string
		Age  int
	}{
		"Person",
		v.Name,
		v.Age,
	}

	b2, _ := json.Marshal(x)
	w.Header().Set("CONTENT-TYPE", "application/json; chartset=utf-8")
	w.Write(b2)

}
