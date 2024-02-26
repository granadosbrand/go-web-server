package database

import (
	"log"
	"strconv"

	"github.com/granadosbrand/go-web-server/api"
)

func (db *DB) UpdateUser(id string, params api.BodyParams) (User, error) {
	userId, err := strconv.Atoi(id)
	if err != nil {
		return User{}, err
	}

	err = db.EnsureDB()
	if err != nil {
		return User{}, err
	}

	dbData, err := db.LoadDB()
	if err != nil {
		return User{}, err
	}

	user := dbData.Users[userId]

	if params.Email != "" {
		user.Email = params.Email
	}
	if params.Password != "" {
		user.Password = params.Password
	}

	dbData.Users[userId] =user

	// Insertarlo en DB

	err = db.writeDB(dbData)
	if err != nil {
		log.Printf("Error writing Database")
		return User{}, err
	}

	return user, nil

}
