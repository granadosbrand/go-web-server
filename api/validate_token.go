package api

import (
    "fmt"

    "github.com/golang-jwt/jwt/v5"
)



func ValidateToken(tokenString, secret string) (string, error) {
    claims := Claims{}

    token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
        return []byte(secret), nil
    })

    if err != nil {
        return "", err
    }

    if !token.Valid {
        return "", fmt.Errorf("Token inválido")
    }

    fmt.Println("Token válido: ", token.Valid)
    subjectId,_ := token.Claims.GetSubject()
    fmt.Println("Token válido: ", token.Valid)



    return subjectId, nil
}
