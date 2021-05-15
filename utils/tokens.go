package utils

import (
	"errors"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

const ACCESS_TOKEN_EXP_MINUTE time.Duration = time.Minute * 2
const REFRESH_TOKEN_EXP_HOUR time.Duration = time.Hour * 360 // 15 days

// kind is either "access" or "refresh"
func NewToken(userId int, kind string) (*jwt.Token, error) {
	switch kind {
	case "access":
		claims := jwt.MapClaims{
			"authorized": true,
			"user_id":    userId,
			"exp":        time.Now().Add(ACCESS_TOKEN_EXP_MINUTE).Unix(),
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		return token, nil
	case "refresh":
		claims := jwt.MapClaims{
			"authorized": true,
			"user_id":    userId,
			"exp":        time.Now().Add(REFRESH_TOKEN_EXP_HOUR).Unix(),
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		return token, nil
	}
	return nil, errors.New("invalid kind")
}
