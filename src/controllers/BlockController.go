package controllers

import (
	"fmt"
	"net/http"
)

func findBlock (w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", r.URL.Path)
}

func BlockController () {
	http.HandleFunc("/block/list", findBlock)
}