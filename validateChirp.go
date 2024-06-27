package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type parameters struct {
	Body string
}

type errorResponse struct {
	Error string
}

type validResponse struct {
	Valid bool `json:"valid"`
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) error {
	response, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(code)
	w.Write(response)
	return nil
}

// func respondWithError(w http.ResponseWriter, code int, msg string) error {
// 	return respondWithJson(w, code, map[string]string{"error": msg})
// }

func handlerValidateChirp(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		log.Printf("Error decoding parameters: %s", err)
		w.WriteHeader(500)
	}

	messageLen := len(params.Body)

	if messageLen > 140 {
		lenErr := errorResponse{
			Error: "Chirp is too long",
		}

		dat, err := json.Marshal(lenErr)
		if err != nil {
			marshalErr := errorResponse{
				Error: "Something went wrong",
			}

			_ = respondWithJson(w, 200, marshalErr)

		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		w.Write(dat)

		return
	}

	validRes := validResponse{
		Valid: true,
	}

	err = respondWithJson(w, 200, validRes)
	if err != nil {
		marshalErr := errorResponse{
			Error: "Something went wrong",
		}

		respondWithJson(w, 200, marshalErr)
	}

}
