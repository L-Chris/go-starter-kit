package main

import (
	"net/http"
	"log"
	"./controllers"
)


func main () {
	controllers.UserController()
	controllers.LayoutController()
	controllers.BlockController()
	controllers.WidgetController()

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}