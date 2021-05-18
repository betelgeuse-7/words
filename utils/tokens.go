package utils

import (
	"errors"
	"fmt"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
)

const ACCESS_TOKEN_EXP_MINUTE time.Duration = time.Minute * 2
const REFRESH_TOKEN_EXP_HOUR time.Duration = time.Hour * 360 // 15 days

// kind is either "access" or "refresh"
func NewToken(userId int, kind string) (string, error) {
	if err := godotenv.Load(); err != nil {
		return "", err
	}
	ACCESS_SECRET := []byte(os.Getenv("ACCESS_TOKEN_SECRET"))
	REFRESH_SECRET := []byte(os.Getenv("REFRESH_TOKEN_SECRET"))

	switch kind {
	case "access":
		claims := jwt.MapClaims{
			"authorized": true,
			"user_id":    userId,
			"exp":        time.Now().Add(ACCESS_TOKEN_EXP_MINUTE).Unix(),
		}
		token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(ACCESS_SECRET)
		if err != nil {
			fmt.Println("access :", err)
			return "", err
		}
		return token, nil
	case "refresh":
		claims := jwt.MapClaims{
			"authorized": true,
			"user_id":    userId,
			"exp":        time.Now().Add(REFRESH_TOKEN_EXP_HOUR).Unix(),
		}
		token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(REFRESH_SECRET)
		if err != nil {
			fmt.Println("refresh :", err)
			return "", err
		}
		return token, nil
	}
	return "", errors.New("invalid kind")
}
