package controllers

import (
	"fmt"
	"net/http"

	"github.com/betelgeuse-7/words/constants"
	//
	"github.com/julienschmidt/httprouter"
)

func SendTokenPair(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	cookies := r.Cookies()

	var oldAccessToken, oldRefreshToken string

	// Assuming the frontend guy stored cookies by the name ACCESS_TOKEN and REFRESH_TOKEN
	// He is me, actually.
	for _, v := range cookies {
		if v.Name == constants.ACCESS_TOKEN_COOKIE_NAME {
			oldAccessToken = v.Value
		}
		if v.Name == constants.REFRESH_TOKEN_COOKIE_NAME {
			oldRefreshToken = v.Value
		}
	}

	fmt.Printf("oldAccessToken: %v\noldRefreshToken: %v \n", oldAccessToken, oldRefreshToken)
	// TODO ...
}
