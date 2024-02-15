package database

import (
	"log"
)

// CreateChirp creates a new chirp and saves it to disk
func (db *DB) CreateUser(email string) (User, error) {

	err := db.EnsureDB()
	if err != nil {
		return User{}, err
	}

	dbData, err := db.LoadDB()
	if err != nil {
		return User{}, err
	}

	idKey := len(dbData.Users) + 1

	userData := User{
		Email: email,
		Id:   idKey,
	}

	dbData.Users[idKey] = userData

	log.Print("Usuario agregado: ", dbData.Users[idKey])

	err = db.writeDB(dbData)
	if err != nil {
		return User{}, err
	}

	return userData, nil
}
