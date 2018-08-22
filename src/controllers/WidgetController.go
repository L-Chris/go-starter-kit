package controllers

import (
	"fmt"
	"net/http"
)

func findWidget (w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", r.URL.Path)
}

func WidgetController () {
	http.HandleFunc("/widget/list", findWidget)
}