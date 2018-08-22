package controllers

import (
	"fmt"
	"net/http"
)

func findUser (w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", r.URL.Path)
}

func findOneUser (w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "test %q", r.URL.Path)
}

func UserController () {
	http.HandleFunc("/user/list", findUser)
	http.HandleFunc("/user/get", findOneUser)
}