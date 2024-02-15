package database

import (
	"encoding/json"
	"log"
	"os"
)

func (db *DB) writeDB(dbStructure DBStructure) error {

	db.mux.Lock()
	defer db.mux.Unlock()
	
	log.Print("Escribiendo base de datos...")

	data, err := json.Marshal(dbStructure)
	if err != nil {
		log.Print("Error unmarshaling in writeDB: ", err)
		return err
	}
	
	os.WriteFile(db.path, data, 6444)
	
	return nil
}
