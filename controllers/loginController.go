package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/betelgeuse-7/words/models"
	"github.com/betelgeuse-7/words/responses"
	"github.com/betelgeuse-7/words/utils"
	"github.com/julienschmidt/httprouter"
	"golang.org/x/crypto/bcrypt"
)

type loginCred struct {
	Email, Password string
}

/*
* check login credentials' validity.
* return 0 for no error, 1 for MISSING_CREDENTIALS, and 2 for EMAIL_INVALID.
 */
func (lc loginCred) Validate() uint {
	if err := utils.LenGreaterThanZero(lc.Email, lc.Password); err != nil {
		return 1
	}
	if err := utils.ValidateEmail(lc.Email); err != nil {
		return 2
	}
	return 0
}

func Login(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	wantHeader := map[string]string{
		"Content-Type": "application/json",
	}
	if res, msg := utils.CheckRequestHeader(wantHeader, r.Header); !res && msg != "ok" {
		responses.CHECK_HEADER_FAIL.Send(w)
		return
	}
	lc := loginCred{}

	json.NewDecoder(r.Body).Decode(&lc)

	validateErrorCode := lc.Validate()
	switch validateErrorCode {
	case 1:
		responses.MISSING_CREDENTIALS.Send(w)
		return
	case 2:
		responses.EMAIL_INVALID.Send(w)
		return
	}

	creds, err := models.GetUserCredsByEmail(lc.Email)
	if err != nil { // there's an internal error (db)
		w.WriteHeader(500)
		responses.LOGIN_FAIL.Send(w)
		return
	}

	// there's no such user record in the db
	if creds.UserId == 0 {
		responses.LOGIN_FAIL.Send(w)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(creds.Password), []byte(lc.Password)); err != nil {
		responses.LOGIN_FAIL.Send(w)
		return
	}

	refreshToken, _ := utils.NewToken(int(creds.UserId), "refresh")
	accessToken, _ := utils.NewToken(int(creds.UserId), "access")

	utils.JSON(w, map[string]string{
		"refresh_token": refreshToken,
		"access_token":  accessToken,
	})
}
