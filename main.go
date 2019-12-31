package main

import (
	"github.com/gorilla/mux"
	userController "go-authen/controllers/user"
	"go-authen/middlewares"
	"go-authen/repo"
	"log"
	"net/http"
	"os"
)

func main() {
	router := mux.NewRouter()

	userPrefix := router.PathPrefix("/userController").Subrouter()
	userPrefix.Path("/new").HandlerFunc(userController.CreateUser).Methods("POST")
	userPrefix.Path("/login").HandlerFunc(userController.LoginUser).Methods("POST")
	userPrefix.Path("/me").HandlerFunc(middlewares.JwtAuthenticate(userController.GetMe)).Methods("GET")

	port := os.Getenv("PORT") // port in which the app will run on
	if port == "" {
		port = "8000"
	}
	log.Println("Application running on port " + port)

	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		log.Panic(err)
	}

	defer repo.GetDB().Close() // close DB connection when exiting
}
