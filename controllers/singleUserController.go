package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/betelgeuse-7/words/models"
	"github.com/betelgeuse-7/words/responses"
	"github.com/betelgeuse-7/words/utils"
	"github.com/julienschmidt/httprouter"
)

func SingleUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := strconv.ParseInt(ps.ByName("id"), 32, 32)
	if err != nil {
		log.Println(err)
		return
	}
	user, err := models.GetSingleUser(int(id))

	if err != nil {
		log.Println("single user err: ", err)
		responses.SERVER_ERROR.Send(w)
		return
	}

	utils.JSON(w, user)
}
