package controllers

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/betelgeuse-7/words/responses"
	"github.com/betelgeuse-7/words/utils"
	"github.com/dgrijalva/jwt-go"

	"github.com/julienschmidt/httprouter"
)

type TToken string

// decode token's payload segment and return it as a []byte
func (t TToken) getPayload() ([]byte, error) {
	return jwt.DecodeSegment(utils.GetTokenPayload(string(t)))
}

func (t TToken) getExpiry() (time.Time, error) {
	return utils.ConvertStringToUnix(utils.GetExpTimeFromTokenPayload([]byte(t)))
}

// ensure the token expired not more than 30 seconds ago.
// return nil if check succeeds.
func (t TToken) checkExpired() error {
	exp, err := t.getExpiry()
	if err != nil {
		return err
	}
	if secondsPassed := time.Since(exp).Seconds(); secondsPassed > 30 {
		return errors.New("token expired more than 30 seconds ago (#) ")
	}
	return nil
}

func (t TToken) getUserId() (int, error) {
	userId, err := utils.GetUserIdFromTokenPayload([]byte(t))
	if err != nil {
		return 0, err
	}
	return userId, nil
}

// send a new refresh and access token pair
func SendTokenPair(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	tokens, err := utils.GetTokensFromCookies(r)
	if err != nil {
		w.WriteHeader(401)
		return
	}
	var oldAccessToken, oldRefreshToken TToken = TToken(tokens[0]), TToken(tokens[1])

	refreshToken, err := jwt.Parse(string(oldRefreshToken), utils.GetSecret)
	if err != nil {
		responses.TOKEN_ERROR.Send(w)
		return
	}

	// ** refresh_token is valid
	if refreshToken.Valid {
		// * check if the user is eligible to get back a token pair
		refreshTokenPayload, err1 := oldRefreshToken.getPayload()
		accessTokenPayload, err2 := oldAccessToken.getPayload()
		if err1 != nil || err2 != nil {
			log.Println("err1: ", err1, "err2: ", err2)
		}
		err = TToken(accessTokenPayload).checkExpired()
		if err != nil {
			w.WriteHeader(401)
			responses.ACCESS_TOKEN_TOO_OLD.Send(w)
			return
		}

		refreshTokenUserId, err1 := TToken(refreshTokenPayload).getUserId()
		accessTokenUserId, err2 := TToken(accessTokenPayload).getUserId()
		if err1 != nil || err2 != nil {
			fmt.Println("errrr", err1, err2)
			w.WriteHeader(500)
			return
		}

		// * check user_id s are the same (for both tokens)
		if refreshTokenUserId != accessTokenUserId {
			w.WriteHeader(401)
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
