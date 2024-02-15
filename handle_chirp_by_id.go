package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/granadosbrand/go-web-server/intern/database"
)

func handleChirpByID(w http.ResponseWriter, r *http.Request) {

	param := chi.URLParam(r, "chirpId")
	chirpID, err := strconv.Atoi(param)
	log.Print("ID convertido: ", chirpID)
	if err != nil {
		w.WriteHeader(400)
		log.Print("Error reading ID: ", err)
		return
	}


	db, err := database.NewDB("./database.json")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)             
		jsonResponse := `{"error": "Id number not valid"}` 
		w.Write([]byte(jsonResponse))                 
		return
	}

	err = db.EnsureDB()
	if err != nil {
		w.WriteHeader(400)
		log.Print("Error reading database: ", err)
		return
	}

	dbData, err := db.LoadDB()
	if err != nil {
		w.WriteHeader(400)
		log.Print("Error reading database: ", err)
		return
	}

	requiredChirp, ok := dbData.Chirps[chirpID]
	if !ok {
		w.WriteHeader(http.StatusNotFound)             
		jsonResponse := `{"error": "Chirp not found"}` 
		w.Write([]byte(jsonResponse))                  
		return
	}

	res, err := json.Marshal(requiredChirp)
	if err != nil {
		log.Print("Could not Marshall the Chirp")
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(res))

}
