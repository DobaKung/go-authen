package controllers

import (
	"encoding/json"
	"go-authen/repo"
	"go-authen/utils"
	"log"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	// Decode request into struct
	usr := &repo.User{}
	decErr := json.NewDecoder(r.Body).Decode(usr)
	if decErr != nil {
		log.Println("CreateUser: decoding error")
		resp := utils.NewMessage(false, 1, "wrong message format", nil)
		utils.Respond(w, http.StatusBadRequest, resp)
		return
	}

	// Create a user
	if err := usr.Create(); err != nil {
		log.Println("CreateUser: " + err.Error())
		resp := utils.NewMessage(false, 2, err.Error(), nil)
		utils.Respond(w, http.StatusConflict, resp)
		return
	}

	// Creation successful, send a response
	accountPayload := repo.UserPayload{
		ID:       usr.ID,
		FullName: usr.FullName,
		Email:    usr.Email,
	}
	response := utils.NewMessage(true, 0, "a user is created", accountPayload)
	utils.Respond(w, http.StatusCreated, response)
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	// Decode request into struct
	user := &repo.User{}
	decErr := json.NewDecoder(r.Body).Decode(user)
	if decErr != nil {
		log.Println("LoginUser: decoding error")
		resp := utils.NewMessage(false, 1, "wrong message format", nil)
		utils.Respond(w, http.StatusBadRequest, resp)
		return
	}

	// Log in
	token, err := repo.Login(user.Email, user.Password)

	// Respond with the correct error code
	if err != nil {
		log.Println("LoginUser: " + err.Error())
		var errCode = 3
		if err == repo.ErrInvalidLoginCred {
			errCode = 4
		}
		resp := utils.NewMessage(false, errCode, err.Error(), nil)
		utils.Respond(w, http.StatusUnauthorized, resp)
		return
	}

	// Respond with token
	tokenPayload := map[string]interface{}{"token": token}
	response := utils.NewMessage(true, 0, "logged in", tokenPayload)
	utils.Respond(w, http.StatusOK, response)
}

func GetMe(w http.ResponseWriter, r *http.Request) {
	response := utils.NewMessage(true, 0, "this is me", nil)
	utils.Respond(w, http.StatusOK, response)
}