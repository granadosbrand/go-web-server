package database

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

// CreateChirp creates a new chirp and saves it to disk
func (db *DB) GetUserLogin(email, password string) (User, error) {

	err := db.EnsureDB()
	if err != nil {
		return User{}, err
	}

	dbData, err := db.LoadDB()
	if err != nil {
		return User{}, err
	}

	usersMap := dbData.Users
	userAthenticated := User{}

	for _, user := range usersMap {
		if user.Email == email {
			userAthenticated = user
		}
	}

	if userAthenticated.Password == "" {
		return User{}, errors.New("User doesn´t exists")
	}

	err = bcrypt.CompareHashAndPassword([]byte(userAthenticated.Password), []byte(password))
	if err != nil {
		return User{}, err
	}

	return userAthenticated, nil
}
