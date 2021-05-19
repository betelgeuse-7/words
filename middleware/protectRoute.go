package middleware

import (
	"fmt"
	"net/http"

	"github.com/betelgeuse-7/words/constants"
	"github.com/betelgeuse-7/words/utils"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/julienschmidt/httprouter"
)

func ProtectRoute(endpoint func(w http.ResponseWriter, r *http.Request, ps httprouter.Params)) httprouter.Handle {
	var handlerFunc = func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		cookies := r.Cookies()

		for _, v := range cookies {
			if v.Name == constants.ACCESS_TOKEN_COOKIE_NAME {
				if v.Value != "" {
					token, err := jwt.Parse(v.Value, utils.GetAccessSecret)
					if err != nil {
						fmt.Println("PROTECT ROUTE: ", err)
						w.WriteHeader(401)
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
