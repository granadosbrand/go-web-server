package api

import (
	"fmt"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)



func GenerateToken(exp int, userID int, secret string) (string, error) {


	secretKey := []byte(secret)

	fmt.Println("exp: ", exp)
	fmt.Println("userUD: ", userID)
	fmt.Println("secret: ", secret)

	expireTime := func () time.Time {
			if exp == 0 {
				return time.Now().Add(500 * time.Hour)
			} else {
				return time.Now().Add(time.Duration(exp) * time.Hour)
			}
	}

	claims := Claims{
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expireTime()),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "chirpy",
			Subject:   strconv.Itoa(userID),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(secretKey)

	return ss, err

}
