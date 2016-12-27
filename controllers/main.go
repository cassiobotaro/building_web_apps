package main

import (
	"net/http"

	render "gopkg.in/unrolled/render.v1"
)

// Action defines a standard function signature for us to use when creating
// controller actions. A controller action is basically just a method attached to
// a controller.
type Action func(rw http.ResponseWriter, r *http.Request) error

// This is our Base Controller
type AppController struct{}

// The action function helps with error handling in a controller
func (c *AppController) Action(a Action) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		if err := a(rw, r); err != nil {
			http.Error(rw, err.Error(), 500)
		}
	})
}

type MyController struct {
	AppController
	*render.Render
}

func (c *MyController) Index(rw http.ResponseWriter, r *http.Request) error {
	c.JSON(rw, 200, map[string]string{"Hello": "JSON"})
	return nil
}

func (c *MyController) Detail(rw http.ResponseWriter, r *http.Request) error {
	c.JSON(rw, 200, map[string]string{"Detail": "detail"})
	return nil
}

// The action function helps with error handling in a controller
func (c *MyController) Action(a Action) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		if err := a(rw, r); err != nil {
			c.HTML(rw, http.StatusInternalServerError, err.Error(), nil)
		}
	})
}

func main() {
	mux := http.NewServeMux()
	c := &MyController{Render: render.New(render.Options{})}
	mux.Handle("/", c.Action(c.Index))
	mux.Handle("/detail", c.Action(c.Detail))

	http.ListenAndServe(":8080", mux)
}
