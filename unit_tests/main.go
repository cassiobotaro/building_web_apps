package main

import (
	"fmt"
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

func HelloWorld(res http.ResponseWriter, req *http.Request) {
	fmt.Fprint(res, "Hello World")
}

func PostSomething(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "post something")
}

func main() {
	n := negroni.Classic()
	router := mux.NewRouter()
	router.HandleFunc("/", HelloWorld).Methods(http.MethodGet)
	router.HandleFunc("/", PostSomething).Methods(http.MethodPost)
	n.UseHandler(router)
	http.ListenAndServe(":3000", n)
}
