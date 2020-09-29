package api

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"go_banking/helpers"
	"go_banking/users"
	"io/ioutil"
	"log"
	"net/http"
)

type Login struct {
	Username string
	Password string
}

type ErrResponse struct {
	Message string
}

func login(w http.ResponseWriter, r *http.Request) {
	// читаем боди
	body, err := ioutil.ReadAll(r.Body)
	helpers.HandleErr(err)

	// обрабатываем логин
	var formattedBody Login
	err = json.Unmarshal(body, &formattedBody)
	helpers.HandleErr(err)
	login := users.Login(formattedBody.Username, formattedBody.Password)
	fmt.Println("login===>", login)

	// подготавливаем ответ
	if login["message"] == "all is fine" {
		resp := login
		// возвращаем json
		json.NewEncoder(w).Encode(resp)
	} else {
		resp := ErrResponse{Message: "Wrong username or password"}
		// возвращаем json
		json.NewEncoder(w).Encode(resp)
	}



}

func StartApi () {
	router := mux.NewRouter()
	router.HandleFunc("/login", login).Methods("POST")
	fmt.Println("App is working on port :8888")
	log.Fatal(http.ListenAndServe(":8888", router))

}