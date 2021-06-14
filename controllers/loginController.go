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

/*
type Credential interface {
	Validate() error
}
*/
type loginCred struct {
	Email, Password string
}

/**
errType Codes:
	0: OK
	1: MISSING_CREDENTIALS
	2: EMAIL_INVALID
	3: LOGIN_FAIL
*/
/*
func (lc *loginCred) Validate() int {
	var errType int

	if err := utils.LenGreaterThanZero(lc.Email, lc.Password); err != nil {
		errType = 1
	}
	if err := utils.ValidateEmail(lc.Email); err != nil {
		errType = 2
	}

	errType = 0
	return errType
}
*/
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

	if err := utils.LenGreaterThanZero(lc.Email, lc.Password); err != nil {
		responses.MISSING_CREDENTIALS.Send(w)
		return
	}
	if err := utils.ValidateEmail(lc.Email); err != nil {
		responses.EMAIL_INVALID.Send(w)
		return
	}

	creds, err := models.GetUserCredsByEmail(lc.Email)
	if err != nil {
		responses.LOGIN_FAIL.Send(w)
		return
	}

	if userId, err := creds.UserId, utils.LenGreaterThanZero(creds.Password); userId < 1 || err != nil {
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
