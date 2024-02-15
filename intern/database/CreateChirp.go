package database

import (
	"log"
)

// CreateChirp creates a new chirp and saves it to disk
func (db *DB) CreateChirp(body string) (Chirp, error) {

	err := db.EnsureDB()
	if err != nil {
		return Chirp{}, err
	}

	dbData, err := db.LoadDB()
	if err != nil {
		return Chirp{}, err
	}

	idKey := len(dbData.Chirps) + 1

	chirpContent := Chirp{
		Body: body,
		Id:   idKey,
	}

	dbData.Chirps[idKey] = chirpContent

	log.Print("Chirp agregado: ", dbData.Chirps)

	err = db.writeDB(dbData)
	if err != nil {
		return Chirp{}, err
	}

	return chirpContent, nil
}
