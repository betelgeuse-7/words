package main

import (
	"log"
	"net/http"

	"github.com/betelgeuse-7/words/controllers"
	"github.com/betelgeuse-7/words/models"

	// Third party
	"github.com/julienschmidt/httprouter"
)

const PORT string = ":8000"

func main() {
	err := models.InitDB()
	if err != nil {
		log.Fatalln(err)
	}

	setup()

	log.Println("Server started on port: ", PORT)
}

func setup() {
	router := httprouter.New()

	router.GET("/api/users", controllers.AllUsers)
	router.GET("/api/user/:id", controllers.SingleUser)

	router.POST("/api/auth/register", models.Register)

	log.Fatalln(http.ListenAndServe(PORT, router))
}
