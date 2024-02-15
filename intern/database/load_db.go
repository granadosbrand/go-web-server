package database

import (
	"encoding/json"
	"log"
	"os"
)

// loadDB reads the database file into memory
func (db *DB) LoadDB() (DBStructure, error) {
	db.mux.RLock()
	defer db.mux.RUnlock()

	dbData := DBStructure{}

	data, err := os.ReadFile(db.path)
	if err != nil {
		log.Print("Couldn't read file")
		return DBStructure{}, err
	}

	err = json.Unmarshal(data, &dbData)

	if err != nil {
		log.Print("Error unmarshaling data into dbData: ", err)
		return DBStructure{}, err
	}

	log.Print("Leyendo DB")
	log.Print("Datos actuales DB: ", dbData.Chirps)

	return dbData, nil
}
