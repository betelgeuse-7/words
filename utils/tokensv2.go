package utils

import (
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
)

type TToken interface {
	/*
	* create a new token
	 */
	MakeNew() (string, error)
}

type AccessToken struct {
	SecretKey []byte
	Payload   struct {
		UserId int
	}
	// expiry date in Unix format
	Expiry int64
}

func (at *AccessToken) MakeNew() (string, error) {
	claims := jwt.MapClaims{
		"authorized": true,
		"user_id":    at.Payload.UserId,
		"exp":        at.Expiry,
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(at.SecretKey)
	if err != nil {
		return "", err
	}
	return token, nil
}

const V2_ACCESS_TOKEN_EXP_MINUTE time.Duration = time.Minute * 2
const V2_REFRESH_TOKEN_EXP_HOUR time.Duration = time.Hour * 360 // 15 days

func NewToken2(userId int) string {
	if err := godotenv.Load(); err != nil {
		return ""
	}
	ACCESS_SECRET := []byte(os.Getenv("ACCESS_TOKEN_SECRET"))
	//REFRESH_SECRET := []byte(os.Getenv("REFRESH_TOKEN_SECRET"))

	atStruct := AccessToken{
		SecretKey: ACCESS_SECRET,
		Payload: struct{ UserId int }{
			UserId: userId,
		},
		Expiry: time.Now().Add(V2_ACCESS_TOKEN_EXP_MINUTE).Unix(),
	}
	accessToken, err := atStruct.MakeNew()
	if err != nil {
		return ""
	}
	return accessToken
}
