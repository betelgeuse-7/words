package utils

import (
	"errors"
	"log"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
)

func GetAccessSecret(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, errors.New("couldn't verify signing method")
	}

	if err := godotenv.Load(); err != nil {
		log.Println(err)
		return nil, err
	}

	ACCESS_SECRET := os.Getenv("ACCESS_TOKEN_SECRET")

	return []byte(ACCESS_SECRET), nil
}

func GetRefreshSecret(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, errors.New("couldn't verify signing method")
	}

	if err := godotenv.Load(); err != nil {
		log.Println(err)
		return nil, err
	}

	REFRESH_SECRET := os.Getenv("REFRESH_TOKEN_SECRET")

	return []byte(REFRESH_SECRET), nil
}
