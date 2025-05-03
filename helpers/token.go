package helpers

import (
	"fmt"
	"golang-jwt/models"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var mySigninKey = []byte("pinjamduluseratus")

type MyCustomcClaims struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	jwt.RegisteredClaims
}

func CreateToken(user *models.User) (string, error) {
	claims := MyCustomcClaims{
		user.ID,
		user.Name,
		user.Email,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(2 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigninKey)

	return ss, err
}

func ValidateToken(tokenString string) (any, error) {
	token, err := jwt.ParseWithClaims(tokenString, &MyCustomcClaims{}, func(t *jwt.Token) (interface{}, error) {
		return mySigninKey, nil
	})

	if err != nil {
		return nil, fmt.Errorf("Unauthorized")
	}

	claims, ok := token.Claims.(*MyCustomcClaims)

	if !ok || !token.Valid {
		return nil, fmt.Errorf("Unauthorized")
	}

	return claims, nil
}
