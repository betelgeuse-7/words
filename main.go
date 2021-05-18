package main

import (
	"log"
	"net/http"

	"github.com/betelgeuse-7/words/controllers"
	"github.com/betelgeuse-7/words/middleware"
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

	log.Println("Server started on port: ", PORT)

	setup()
}

func setup() {
	router := httprouter.New()

	router.GET("/api/users", middleware.IsLoggedIn(controllers.AllUsers))
	router.GET("/api/user/:id", middleware.IsLoggedIn(controllers.SingleUser))

	router.POST("/api/auth/register", models.Register)
	router.POST("/api/auth/login", controllers.Login)
	// Need to be logged in
	router.POST("/api/auth/refresh", middleware.IsLoggedIn(controllers.SendTokenPair))

	log.Fatalln(http.ListenAndServe(PORT, router))
}
