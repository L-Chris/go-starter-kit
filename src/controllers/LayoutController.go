package controllers

import (
	"fmt"
	"net/http"
)

func findLayout (w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", r.URL.Path)
}

func LayoutController () {
	http.HandleFunc("/layout/list", findLayout)
}