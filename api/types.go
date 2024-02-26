package api

import "github.com/golang-jwt/jwt/v5"

type Claims struct {
	jwt.RegisteredClaims
}

type BodyParams struct {
	Password string `json:"password"`
	Email    string `json:"email"`
}
