package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/betelgeuse-7/words/constants"
	"github.com/betelgeuse-7/words/responses"
	"github.com/betelgeuse-7/words/utils"
	"github.com/dgrijalva/jwt-go"

	//
	"github.com/julienschmidt/httprouter"
)

/*
type payload struct {
	Authorized  bool
	Exp, UserId uint
}
*/
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
	// Missing token in cookies
	if err := utils.LenGreaterThanZero(oldAccessToken, oldRefreshToken); err != nil {
		w.WriteHeader(401)
		return
	}

	refreshToken, err := jwt.Parse(oldRefreshToken, utils.GetRefreshSecret)
	if err != nil {
		log.Println("REFRESH1: ", err)
		json.NewEncoder(w).Encode(responses.TOKEN_ERROR)
		return
	}

	if refreshToken.Valid {

		// ** 1) check user_id s are the same (for both tokens)
		// ** 2) give tokens with a user_id the same as that of the access_token
		// ** 3) check access_token expired less than 30 secs ago.

		// if a bad guy only has someones refresh token he will not be able to
		// get access to the victim's account. (1)

		// if we give a new pair of tokens with user_id of a refresh_token.
		// the bad guy will have the access to the victim's account
		// with JUST a refresh token. (2)

		// security :) (3)

		// if the bad guy has only the refresh token. he will not be able to get a
		// new pair back (1)

		// if the bad guy has both the tokens. it is f***ed up.

		// TODO...
		/*
			 ! BOOM
				refreshTokenUserId, err := strconv.ParseInt(utils.GetUserIdFromTokenPayload([]byte(oldRefreshToken)), 10, 64)
				if err != nil {
					fmt.Println(err)
					w.WriteHeader(500)
					return
				}
				accessTokenUserId, err := strconv.ParseInt(utils.GetUserIdFromTokenPayload([]byte(oldAccessToken)), 10, 64)
				if err != nil {
					fmt.Println(err)
					w.WriteHeader(500)
					return
				}

				if refreshTokenUserId == accessTokenUserId {
					fmt.Println("SAME")
					fmt.Println("REFRESH USER ID: ", refreshTokenUserId)
					fmt.Println("ACCESS USER ID: ", accessTokenUserId)
				}
		*/
		fmt.Println("REFRESH TOKEN VALID")
		decodedRefreshTokenPayload, err := jwt.DecodeSegment(strings.Split(oldRefreshToken, ".")[1])
		if err != nil {
			log.Println("REFRESH2: ", err)
		}
		decodedAccessTokenPayload, err := jwt.DecodeSegment(strings.Split(oldAccessToken, ".")[1])
		if err != nil {
			log.Println("REFRESH2: ", err)
		}
		unixRefreshExp, err := utils.ConvertStringToUnix(utils.GetExpTimeFromTokenPayload(decodedRefreshTokenPayload))
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(500)
			return
		}
		unixAccessExp, err := utils.ConvertStringToUnix(utils.GetExpTimeFromTokenPayload(decodedAccessTokenPayload))
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(500)
			return
		}

		fmt.Println("Access Token Unix Exp: ", unixAccessExp)
		fmt.Println("Refresh Token Unix Exp: ", unixRefreshExp)
	}
}
