package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/betelgeuse-7/words/responses"
	"github.com/betelgeuse-7/words/utils"
	"github.com/julienschmidt/httprouter"
)

type loginCred struct {
	Email, Password string
}

func Login(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	wantHeader := map[string]string{
		"Content-Type": "application/json",
	}
	if res, msg := utils.CheckRequestHeader(wantHeader, r.Header); !res && msg != "ok" {
		json.NewEncoder(w).Encode(responses.CHECK_HEADER_FAIL)
		return
	}
	lc := loginCred{}

	json.NewDecoder(r.Body).Decode(&lc)

	fmt.Println("LC: ", lc)
	/*
		! Buggy
		if err := utils.LenGreaterThanZero(lc.Email, lc.Password); err != nil {
			json.NewEncoder(w).Encode(responses.MISSING_CREDENTIALS)
			return
		}

		if err := utils.ValidateEmail(lc.Email); err != nil {
			json.NewEncoder(w).Encode(responses.EMAIL_INVALID)
			return
		}

		creds, err := models.GetUserCredsByEmail(lc.Email)
		if err != nil {
			json.NewEncoder(w).Encode(responses.MISSING_CREDENTIALS)
			return
		}

		if userId, err := creds.UserId, utils.LenGreaterThanZero(creds.Password); userId < 1 || err != nil {
			json.NewEncoder(w).Encode(responses.LOGIN_FAIL)
			return
		}

		fmt.Println(creds)
	*/
}