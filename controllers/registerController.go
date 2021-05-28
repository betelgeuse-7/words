package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/betelgeuse-7/words/models"
	"github.com/betelgeuse-7/words/responses"
	"github.com/betelgeuse-7/words/utils"
	"github.com/julienschmidt/httprouter"
	"golang.org/x/crypto/bcrypt"
)

type newUser struct {
	FirstName, LastName, Email, Password string
}

func RegisterController(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	wantHeader := map[string]string{
		"Content-Type": "application/json",
	}
	if ok, msg := utils.CheckRequestHeader(wantHeader, r.Header); ok && msg == "ok" {
		var new newUser
		if err := json.NewDecoder(r.Body).Decode(&new); err != nil {
			json.NewEncoder(w).Encode(responses.REGISTER_FAIL)
			return
		}
		if err := utils.LenGreaterThanZero(new.FirstName, new.LastName, new.Email, new.Password); err != nil {
			json.NewEncoder(w).Encode(responses.MISSING_CREDENTIALS)
			return
		}
		if err := utils.ValidateEmail(new.Email); err != nil {
			json.NewEncoder(w).Encode(responses.EMAIL_INVALID)
			return
		}

		// * encrypt password
		password, err := bcrypt.GenerateFromPassword([]byte(new.Password), 10)
		if err != nil {
			log.Println(err)
			json.NewEncoder(w).Encode(responses.SERVER_ERROR)
			return
		}
		new.Password = string(password)
		registeredAt := time.Now()

		if err := models.Register(new.FirstName, new.LastName, new.Email, new.Password, registeredAt); err != nil {
			log.Println("Register controller", err)
			json.NewEncoder(w).Encode(responses.SERVER_ERROR)
			return
		}
		json.NewEncoder(w).Encode(responses.REGISTER_SUCCESS)
	} else {
		json.NewEncoder(w).Encode(responses.CHECK_HEADER_FAIL)
		return
	}
}
