package controllers

import (
	"net/http"

	"github.com/betelgeuse-7/words/models"
	"github.com/betelgeuse-7/words/responses"
	"github.com/betelgeuse-7/words/utils"
	"github.com/julienschmidt/httprouter"
)

// return all of the user records in the database in a json format
func AllUsers(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	users, err := models.GetUsers()
	if err != nil {
		responses.SERVER_ERROR.Send(w)
		return
	}
	utils.JSON(w, users)
}
