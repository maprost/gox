package core

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/maprost/gox/testexample/userserver/client"
	"github.com/maprost/gox/testexample/userserver/core/data"
	"net/http"
	"os/user"
)

func InitUser(router *mux.Router) {
	router.Path("/user").Methods(http.MethodPost).Handler(createUser)
	router.Path("/user/{email}").Methods(http.MethodDelete).Handler(deleteUser)
	router.Path("/user").Methods(http.MethodGet).Handler(getUsers)
}

func createUser(writer http.ResponseWriter, request *http.Request) {
	defer request.Body.Close()

	var userDto client.UserDto
	err := json.NewDecoder(request.Body).Decode(&userDto)
	if err != nil {
		http.Error(writer, "Can't open UserDto", http.StatusNotFound)
		return
	}

	user := data.UserAccount{
		Email:        userDto.Email,
		PasswordHash: userDto.PasswordHash,
	}

	err = user.Insert()
	if err != nil {
		http.Error(writer, "Can't store user in database.", http.StatusBadRequest)
		return
	}

	writer.WriteHeader(http.StatusNoContent)
}

func deleteUser(writer http.ResponseWriter, request *http.Request) {
	email, ok := mux.Vars(request)["email"]
	if !ok {
		http.Error(writer, "Can't delete user out of database.", http.StatusBadRequest)
		return
	}

	err := data.DeleteUser(email)
	if err != nil {
		http.Error(writer, "Can't delete user out of database.", http.StatusBadRequest)
		return
	}

	writer.WriteHeader(http.StatusNoContent)
}

func getUsers(writer http.ResponseWriter, request *http.Request) {
	users, err := data.GetUsers()
	if err != nil {
		http.Error(writer, "Can't load user out of database.", http.StatusBadRequest)
		return
	}

	writer.WriteHeader(http.StatusNoContent)
}
