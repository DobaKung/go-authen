package main

import (
	"github.com/gorilla/mux"
	"go-authen/controllers"
	"go-authen/middlewares"
	"go-authen/repo"
	"log"
	"net/http"
	"os"
)

func main() {
	router := mux.NewRouter()

	userPrefix := router.PathPrefix("/user").Subrouter()
	userPrefix.Path("/new").HandlerFunc(controllers.CreateUser).Methods("POST")
	userPrefix.Path("/login").HandlerFunc(controllers.LoginUser).Methods("POST")
	userPrefix.Path("/me").HandlerFunc(middlewares.JwtAuthenticate(controllers.GetMe)).Methods("GET")

	port := os.Getenv("PORT") // port in which the app will run on
	if port == "" {
		port = "8000"
	}
	log.Println("Application running on port " + port)

	err := http.ListenAndServe(":" + port, router)
	if err != nil {
		log.Panic(err)
	}

	defer repo.GetDB().Close() // close DB connection when exiting
}
