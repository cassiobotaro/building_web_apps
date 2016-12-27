package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_HelloWorld(t *testing.T) {
	req, err := http.NewRequest("GET", "http://127.0.0.1/", nil)
	if err != nil {
		t.Fatal(err)
	}

	res := httptest.NewRecorder()
	HelloWorld(res, req)

	exp := "Hello World"
	act := res.Body.String()
	if exp != act {
		t.Fatalf("Expected %s gog %s", exp, act)
	}
}

func Test_PostSomething(t *testing.T) {
	req, err := http.NewRequest("POST", "http://127.0.0.1/", nil)
	if err != nil {
		t.Fatal(err)
	}

	res := httptest.NewRecorder()
	PostSomething(res, req)

	exp := "post something"
	act := res.Body.String()
	if exp != act {
		t.Fatalf("Expected %s gog %s", exp, act)
	}

}
