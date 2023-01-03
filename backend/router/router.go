package router

import (
	"github.com/gorilla/mux"
	"github.com/vishal21121/myapp/controller"
)

func Router() *mux.Router{
	router := mux.NewRouter()
	router.HandleFunc("/createUser",controller.CreateUser).Methods("POST")
	router.HandleFunc("/login",controller.Login).Methods("POST")
	return router
}
