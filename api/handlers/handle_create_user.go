package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/granadosbrand/go-web-server/intern/database"
)

type createUSerParams struct {
	Password string `json:"password"`
	Email    string `json:"email"`
}

func HandleCreateUser(w http.ResponseWriter, r *http.Request) {

	db, err := database.NewDB("./database.json")
	if err != nil {
		log.Print("Error creating DB instance: ", err)
	}

	decoder := json.NewDecoder(r.Body)
	createUSerParams := createUSerParams{}
	err = decoder.Decode(&createUSerParams)
	if err != nil {
		// an error will be thrown if the JSON is invalid or has the wrong types
		// any missing fields will simply have their values in the struct set to their zero value
		log.Printf("Error decoding parameters: %s", err)
		w.WriteHeader(500)
		return
	}

	user, err := db.CreateUser(createUSerParams.Email, createUSerParams.Password)
	if err != nil {
		w.WriteHeader(400)
		log.Print("Error creating User")
		return
	}

	resUSer := struct {
		Id int `json:"id"`
		Email    string `json:"email"`
	}{
		Id: user.Id,
		Email: user.Email,
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
