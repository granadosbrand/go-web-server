package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/granadosbrand/go-web-server/intern/database"
)

func HandleChirpGet(w http.ResponseWriter, r *http.Request) {

	db, err := database.NewDB("./database.json")
	if err != nil {
		log.Print("Error creating DB instance: ", err)
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

	resStructData:= []database.Chirp{}

	for i := 1; i < len(dbData.Chirps) + 1; i++ {

		resStructData = append(resStructData, dbData.Chirps[i] )
		
	}

	marshalData, err := json.Marshal(resStructData)
	if err != nil {
		log.Print("Error marshalling DB: ", err)
		w.WriteHeader(400)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(marshalData)

}
