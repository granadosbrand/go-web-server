package database

import (
	"sync"
)

func NewDB(path string) (*DB, error) {

	db := DB{
		path: path,
		mux:  &sync.RWMutex{},
	}

	return &db, nil
}
