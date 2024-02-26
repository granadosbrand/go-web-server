package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/granadosbrand/go-web-server/api"
	"github.com/granadosbrand/go-web-server/config"
	"github.com/granadosbrand/go-web-server/intern/database"
)


func HandleUpdateUser(c *config.ApiConfig, w http.ResponseWriter, r *http.Request) {

	token := strings.Split(r.Header.Get("Authorization"), " ")[1]
	log.Printf("Token recibido: %s", token)
	
	userId, err := api.ValidateToken(token, c.JwtSecret)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		log.Printf("Error validating JWT: %s", err)
		w.Write([]byte("Error validating JWT"))
		return
	}


	decoder := json.NewDecoder(r.Body)
	loginParams := api.BodyParams{}
	err = decoder.Decode(&loginParams)
	if err != nil {
		// an error will be thrown if the JSON is invalid or has the wrong types
		// any missing fields will simply have their values in the struct set to their zero value
		log.Printf("Error decoding parameters: %s", err)
		w.WriteHeader(500)
		return
	}

	db, err := database.NewDB("./database.json")
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Error reading database"))
		return
	}

	updatedUser, err := db.UpdateUser(userId, loginParams)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Error updating user"))
		return
	}

	resUSer := struct {
		Email string `json:"email"`
		Id    int    `json:"id"`
	}{
		Email: updatedUser.Email,
		Id:    updatedUser.Id,
	}

	resData, err := json.Marshal(resUSer)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte(fmt.Sprint("Error marshaling user data: ", err)))
		return
	}

	w.WriteHeader(201)
	w.Write(resData)

}
