package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/betelgeuse-7/words/models"
	"github.com/betelgeuse-7/words/responses"
	"github.com/julienschmidt/httprouter"
)

func AllUsers(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	users, err := models.GetUsers()
	if err != nil {
		log.Println(err, " <|> AllUsers")
		json.NewEncoder(w).Encode(responses.SERVER_ERROR)
		return
	}
	json.NewEncoder(w).Encode(users)
}

func SingleUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := strconv.ParseInt(ps.ByName("id"), 32, 32)
	if err != nil {
		log.Println(err)
		return
	}

	user, err := models.GetSingleUser(int(id))

	if err != nil {
		log.Println(err, "  <|>  SingleUser")
		json.NewEncoder(w).Encode(responses.SERVER_ERROR)
		return
	}

	json.NewEncoder(w).Encode(user)
}
