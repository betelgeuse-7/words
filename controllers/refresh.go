package controllers

import (
	"log"
	"net/http"
	"time"

	"github.com/betelgeuse-7/words/responses"
	"github.com/betelgeuse-7/words/utils"
	"github.com/dgrijalva/jwt-go"

	"github.com/julienschmidt/httprouter"
)

// send a new refresh and access token pair
func SendTokenPair(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var oldAccessToken, oldRefreshToken string

	tokens, err := utils.GetTokensFromCookies(r)
	if err != nil {
		w.WriteHeader(401)
		return
	}
	oldAccessToken, oldRefreshToken = tokens[0], tokens[1]

	refreshToken, err := jwt.Parse(oldRefreshToken, utils.GetRefreshSecret)
	if err != nil {
		responses.TOKEN_ERROR.Send(w)
		return
	}

	// ** refresh_token is valid
	if refreshToken.Valid {
		decodedRefreshTokenPayload, err := jwt.DecodeSegment(utils.GetTokenPayload(oldRefreshToken))
		if err != nil {
			log.Println(err)
		}

		decodedAccessTokenPayload, err := jwt.DecodeSegment(utils.GetTokenPayload(oldAccessToken))
		if err != nil {
			log.Println(err)
		}

		unixAccessExp, err := utils.ConvertStringToUnix(utils.GetExpTimeFromTokenPayload(decodedAccessTokenPayload))
		if err != nil {
			w.WriteHeader(500)
			return
		}

		refreshTokenUserId, err := utils.GetUserIdFromTokenPayload([]byte(decodedRefreshTokenPayload))
		if refreshTokenUserId < 1 || err != nil {
			w.WriteHeader(500)
			return
		}
		accessTokenUserId, err := utils.GetUserIdFromTokenPayload([]byte(decodedAccessTokenPayload))
		if accessTokenUserId < 1 || err != nil {
			w.WriteHeader(500)
			return
		}

		// * check user_id s are the same (for both tokens)
		if refreshTokenUserId != accessTokenUserId {
			w.WriteHeader(401)
			return
		}

		// * check access_token expired less than 30 secs ago
		if secondsPassed := time.Since(unixAccessExp).Seconds(); secondsPassed > 30 {
			w.WriteHeader(401)
			responses.ACCESS_TOKEN_TOO_OLD.Send(w)
			return
		}
		// * give tokens with a user_id the same as that of the access_token
		tokenPair, err := utils.NewTokenPair(int(accessTokenUserId))
		if err != nil {
			w.WriteHeader(500)
			return
		}

		utils.JSON(w, map[string]string{
			"new_refresh_token": tokenPair[1],
			"new_access_token":  tokenPair[0],
		})
	}
}
