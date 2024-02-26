package database

import (
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

// CreateChirp creates a new chirp and saves it to disk
func (db *DB) CreateUser(email, password string) (User, error) {

	err := db.EnsureDB()
	if err != nil {
		return User{}, err
	}

	dbData, err := db.LoadDB()
	if err != nil {
		return User{}, err
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
		return User{}, err
	}

	idKey := len(dbData.Users) + 1

	userData := User{
		Email:    email,
		Password: string(hash),
		Id:       idKey,
	}

	dbData.Users[idKey] = userData

	log.Print("Usuario agregado: ", dbData.Users[idKey])

	err = db.writeDB(dbData)
	if err != nil {
		return User{}, err
	}

	return userData, nil
}
