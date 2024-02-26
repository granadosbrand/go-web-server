package database

import "sync"

type DB struct {
	path string
	mux  *sync.RWMutex
}

type Chirp struct {
	Body string `json:"body"`
	Id   int    `json:"id"`
}
type DBStructure struct {
	Chirps map[int]Chirp `json:"chirps"`
	Users  map[int]User  `json:"users"`
}

type User struct {
	Password string `json:"password"`
	Email string `json:"email"`
	Id    int    `json:"id"`
}
