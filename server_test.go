package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_HelloWorld(t *testing.T) {
	req, err := http.NewRequest("GET", "http://example.com/foo", nil)
	if err != nil {
		t.Fatal(err)
	}

	res := httptest.NewRecorder()
	HelloWorld(res, req, nil)

	exp := "Hello World!\n"
	act := res.Body.String()
	if exp != act {
		t.Fatalf("Expected '%s' got '%s'", exp, act)
	}
}

func Test_App(t *testing.T) {
	ts := httptest.NewServer(App())
	defer ts.Close()

	res, err := http.Get(ts.URL)
	if err != nil {
		t.Fatal(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()

	if err != nil {
		t.Fatal(err)
	}

	exp, err := ioutil.ReadFile("static/index.html")
	if err != nil {
		t.Fatal(err)
	}

	if string(exp) != string(body) {
		t.Fatalf("Expected '%s' got '%s'", exp, body)
	}
}
