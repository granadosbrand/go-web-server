package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"github.com/granadosbrand/go-web-server/config"
	"github.com/granadosbrand/go-web-server/api"
	"github.com/granadosbrand/go-web-server/intern/database"
)

type loginRequest struct {
	Password       string `json:"password"`
	Email          string `json:"email"`
	ExpiresSeconds int    `json:"expires_in_seconds"`
}

func HandleLogin(c *config.ApiConfig, w http.ResponseWriter, r *http.Request) {

	db, err := database.NewDB("./database.json")
	if err != nil {
		log.Print("Error creating DB instance: ", err)
	}

	decoder := json.NewDecoder(r.Body)
	loginParams := loginRequest{}
	err = decoder.Decode(&loginParams)
	if err != nil {
		// an error will be thrown if the JSON is invalid or has the wrong types
		// any missing fields will simply have their values in the struct set to their zero value
		log.Printf("Error decoding parameters: %s", err)
		w.WriteHeader(500)
		return
	}

	userData, err := db.GetUserLogin(loginParams.Email, loginParams.Password)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		log.Print("Error Loging in")
		return
	}

	token, err := api.GenerateToken(loginParams.ExpiresSeconds, userData.Id, c.JwtSecret)
	if err != nil {
		log.Printf("Error generating Token: %s", err)
		w.WriteHeader(500)
		return
	}

	resUSer := struct {
		Id    int    `json:"id"`
		Email string `json:"email"`
		Token string `json:"token"`
	}{
		Id:    userData.Id,
		Email: userData.Email,
		Token: token,
	}

	resData, err := json.Marshal(resUSer)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte(fmt.Sprint("Error marshaling user data: ", err)))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(resData)
}
