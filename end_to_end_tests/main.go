package main

import (
	"fmt"
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/julienschmidt/httprouter"
)

func HelloWorld(res http.ResponseWriter, req *http.Request, p httprouter.Params) {
	fmt.Fprint(res, "Hello World")
}

func HelloPost(res http.ResponseWriter, req *http.Request, p httprouter.Params) {
	fmt.Fprint(res, "Hello "+p.ByName("name"))
}

func MyMiddleware(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	fmt.Fprintf(rw, "Before...")
	next(rw, r)
	fmt.Fprintf(rw, "...After")
}

func AnotherMiddleware(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	fmt.Printf("another middleware\n")
	next(rw, r)
	fmt.Printf("exiting from middleware\n")
}

func App() http.Handler {
	n := negroni.Classic()

	n.UseFunc(MyMiddleware)
	n.UseFunc(AnotherMiddleware)

	r := httprouter.New()

	r.GET("/", HelloWorld)
	r.POST("/:name", HelloPost)
	n.UseHandler(r)
	return n
}

func main() {
	http.ListenAndServe(":3000", App())
}
