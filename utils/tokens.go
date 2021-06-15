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

type TokenClaims struct {
	UserId    uint
	ExpiresAt int64
	jwt.StandardClaims
}

/* will refactor this */

// kind is either "access" or "refresh"
func NewToken(userId int, kind string) (string, error) {
	if err := godotenv.Load(); err != nil {
		return "", err
	}
	//ACCESS_SECRET := []byte(os.Getenv("ACCESS_TOKEN_SECRET"))
	//REFRESH_SECRET := []byte(os.Getenv("REFRESH_TOKEN_SECRET"))
	JWT_SECRET := []byte(os.Getenv("JWT_SECRET"))

	switch kind {
	case "access":
		/*
			claims := jwt.MapClaims{
				"authorized": true,
				"user_id":    userId,
				"exp":        time.Now().Add(ACCESS_TOKEN_EXP_MINUTE).Unix(),
			}*/
		token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, TokenClaims{
			UserId:    uint(userId),
			ExpiresAt: time.Now().Add(ACCESS_TOKEN_EXP_MINUTE).Unix(),
		}).SignedString(JWT_SECRET)
		if err != nil {
			fmt.Println("access :", err)
			return "", err
		}
		return token, nil

	case "refresh":
		/*
			claims := jwt.MapClaims{
				"authorized": true,
				"user_id":    userId,
				"exp":        time.Now().Add(REFRESH_TOKEN_EXP_HOUR).Unix(),
			}*/
		token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, TokenClaims{
			UserId:    uint(userId),
			ExpiresAt: time.Now().Add(REFRESH_TOKEN_EXP_HOUR).Unix(),
		}).SignedString(JWT_SECRET)
		if err != nil {
			fmt.Println("refresh :", err)
			return "", err
		}
		return token, nil
	}
	return "", errors.New("invalid kind")

}

// return new access and refresh tokens
// []string{ACCESS_TOKEN, REFRESH_TOKEN}
func NewTokenPair(userId int) ([]string, error) {
	newAccessToken, err := NewToken(userId, "access")
	if err != nil {
		return nil, err
	}
	newRefreshToken, err := NewToken(userId, "refresh")
	if err != nil {
		return nil, err
	}

	return []string{newAccessToken, newRefreshToken}, nil
}
