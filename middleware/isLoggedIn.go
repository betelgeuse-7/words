package middleware

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/betelgeuse-7/words/constants"
	"github.com/betelgeuse-7/words/responses"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
)

func IsLoggedIn(endpoint func(w http.ResponseWriter, r *http.Request, ps httprouter.Params)) httprouter.Handle {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln(err)
	}
	ACCESS_SECRET := os.Getenv("ACCESS_TOKEN_SECRET")

	var handlerFunc = func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		cookies := r.Cookies()

		getAccessSecret := func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("couldnt verify signing method")
			}
			return []byte(ACCESS_SECRET), nil
		}
		for _, v := range cookies {
			if v.Name == constants.ACCESS_TOKEN_COOKIE_NAME {
				if v.Value != "" {
					token, err := jwt.Parse(v.Value, getAccessSecret)
					if err != nil {
						fmt.Println(err)
						json.NewEncoder(w).Encode(responses.TOKEN_ERROR)
						return
					}
					if token.Valid {
						endpoint(w, r, ps)
					}
				}
			}
		}
	}
	return handlerFunc
}
