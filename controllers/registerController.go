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

func (n *newUser) setPassword(password string) {
	n.Password = password
}

func RegisterController(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	wantHeader := map[string]string{
		"Content-Type": "application/json",
	}
	if ok, msg := utils.CheckRequestHeader(wantHeader, r.Header); ok && msg == "ok" {
		var new newUser

		if err := json.NewDecoder(r.Body).Decode(&new); err != nil {
			responses.REGISTER_FAIL.Send(w)
			return
		}
		if err := utils.LenGreaterThanZero(new.FirstName, new.LastName, new.Email, new.Password); err != nil {
			responses.MISSING_CREDENTIALS.Send(w)
			return
		}
		if err := utils.ValidateEmail(new.Email); err != nil {
			responses.EMAIL_INVALID.Send(w)
			return
		}

		// * encrypt password
		password, err := bcrypt.GenerateFromPassword([]byte(new.Password), 10)
		if err != nil {
			responses.SERVER_ERROR.Send(w)
			return
		}
		new.setPassword(string(password))
		registeredAt := time.Now()

		if err := models.Register(new.FirstName, new.LastName, new.Email, new.Password, registeredAt); err != nil {
			log.Println("Register controller", err)
			responses.SERVER_ERROR.Send(w)
			return
		}
		responses.REGISTER_SUCCESS.Send(w)
	} else {
		responses.CHECK_HEADER_FAIL.Send(w)
		return
	}
}
