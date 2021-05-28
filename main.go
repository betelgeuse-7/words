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

	router.GET("/api/users", middleware.ProtectRoute(controllers.AllUsers))
	router.GET("/api/user/:id", middleware.ProtectRoute(controllers.SingleUser))
	router.GET("/api/my_profile", middleware.ProtectRoute(controllers.MyProfile))

	router.POST("/api/auth/register", controllers.RegisterController)
	router.POST("/api/auth/login", controllers.Login)
	router.POST("/api/auth/refresh", controllers.SendTokenPair)

	log.Fatalln(http.ListenAndServe(PORT, router))
}
