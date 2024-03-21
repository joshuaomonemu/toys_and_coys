package utils

import (
	"app/structs"
	"fmt"
	"github.com/golang-jwt/jwt"
	"log"
	"os"
	"time"
)

func CreateToken(payload structs.JWTPayload) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": payload.UserId,
		"nbf":    time.Now().Add(30 * time.Minute),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(os.Getenv("JWT_SECRET"))

	if err != nil {
		return "", err
	} else {
		return tokenString, nil
	}

}

func VerifyToken(tokenString string) (*structs.JWTPayload, error) {
	decoded_token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		tokenSecret := []byte(os.Getenv("JWT_SECRET"))

		return tokenSecret, nil
	})

	if err != nil {
		log.Fatal(err)
	}

	if claims, ok := decoded_token.Claims.(jwt.MapClaims); ok {
		userId := fmt.Sprintln(claims["userId"])
		return &structs.JWTPayload{
			UserId: userId,
		}, nil
	} else {
		return nil, err
	}
}
