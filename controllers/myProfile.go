package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/betelgeuse-7/words/models"
	"github.com/betelgeuse-7/words/utils"
	"github.com/julienschmidt/httprouter"
)

// ! see models/profile.go#GetProfile
func MyProfile(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var userId int
	var accessToken, accessTokenPayload string

	tokens, err := utils.GetTokensFromCookies(r)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(401)
		return
	}
	accessToken = tokens[0]
	accessTokenPayload = utils.GetTokenPayload(accessToken)
	fmt.Println("ACCESS TOKEN PAYLOAD", accessTokenPayload)
	userId, err = utils.GetUserIdFromTokenPayload([]byte(accessTokenPayload))
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(500)
		return
	}

	profile, err := models.GetProfile(userId)
	if err != nil {
		fmt.Println("MyProfile ERR: ", err)
		w.WriteHeader(500)
		return
	}
	json.NewEncoder(w).Encode(profile)
}
