package database

import (
	"encoding/json"
	"log"
	"os"
)

func NewDBStructure() *DBStructure {
    return &DBStructure{
        Chirps: make(map[int]Chirp),
        Users: make(map[int]User),
    }
}

func (db *DB) EnsureDB() error {
    dbStruct := NewDBStructure()

    jsonData, err := json.Marshal(dbStruct)
    if err != nil {
        log.Print("Error marshaling dbStruct: ", err)
        return err
    }

    db.mux.RLock()
    defer db.mux.RUnlock()

    _, err = os.ReadFile(db.path)
    if err != nil {
        if os.IsNotExist(err) {
            log.Print("Archivo no encontrado, creando uno nuevo: ", err)
            err = os.WriteFile(db.path, jsonData, 0666)
            if err != nil {
                log.Print("Error escribiendo en el archivo: ", err)
                return err
            }
        } else {
            log.Print("Error leyendo el archivo: ", err)
            return err
        }
    }

    return nil
}
