package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"slices"
	"strings"

	"github.com/granadosbrand/go-web-server/api"
	"github.com/granadosbrand/go-web-server/intern/database"
)

type parameters struct {
	Body string `json:"body"`
}


func HandleChirpPost( w http.ResponseWriter, r *http.Request) {

	db , err := database.NewDB("./database.json")
	if err != nil {
		log.Print("Error creating DB instance: ", err)
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err = decoder.Decode(&params)
	if err != nil {
		// an error will be thrown if the JSON is invalid or has the wrong types
		// any missing fields will simply have their values in the struct set to their zero value
		log.Printf("Error decoding parameters: %s", err)
		w.WriteHeader(500)
		return
	}

	if len(params.Body) > 140 {
		api.RespondWithError(w, 400, "Chirp is too long")
		return
	}

	cleaned_chirp := handleValidateChirp(w, params.Body)

	chirp, err := db.CreateChirp(cleaned_chirp)
	if err != nil {
		w.WriteHeader(400)
		log.Print("Error creating Chirp")
		return
	}
	
	resData, err := json.Marshal(chirp)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte(fmt.Sprint("Error marshaling data: ", err)))
		return
	}

	w.WriteHeader(201)
	w.Write(resData)
	

}

func handleValidateChirp(w http.ResponseWriter, chirp string) string {

	// Len validation

	// Profane Words cleaning
	profane_words := []string{"kerfuffle", "sharbert", "fornax"}

	words := strings.Split(chirp, " ")

	for i := 0; i < len(words); i++ {
		if slices.Contains(profane_words, strings.ToLower(words[i])) {
			words[i] = "****"
		}
	}

	cleanedBody := strings.Join(words, " ")

	return cleanedBody

}
