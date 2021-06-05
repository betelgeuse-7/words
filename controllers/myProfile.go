package controllers

import (
	"net/http"

	"github.com/betelgeuse-7/words/models"
	"github.com/betelgeuse-7/words/utils"
	"github.com/julienschmidt/httprouter"
)

func MyProfile(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var userId int
	var accessToken, accessTokenPayload string

	tokens, err := utils.GetTokensFromCookies(r)
	if err != nil {
		w.WriteHeader(401)
		return
	}
	accessToken = tokens[0]
	accessTokenPayload = utils.GetTokenPayload(accessToken)
	userId, err = utils.GetUserIdFromTokenPayload([]byte(accessTokenPayload))
	if err != nil {
		w.WriteHeader(500)
		return
	}

	profile, err := models.GetProfile(userId)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	utils.JSON(w, profile)
}
