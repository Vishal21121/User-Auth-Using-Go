package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/vishal21121/myapp/controller"
	"github.com/vishal21121/myapp/router"
)

func main() {
	r := router.Router()
	fmt.Println("Server is starting")
	controller.Init()
	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("Listening at port 4000")
}
