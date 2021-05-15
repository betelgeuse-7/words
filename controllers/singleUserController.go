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
