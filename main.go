package main

import (
	"github.com/joho/godotenv"
	"gonosql/internal/config"
	"gonosql/internal/handler"
	"log"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	errConnectEnv := godotenv.Load()

	if errConnectEnv != nil {
		log.Fatal("Error loading .env file")
	}
	config.ConnectMongo()
	config.ConnectMysql()

	// Init Controller
	feedUserController := handler.FeedUserController(config.MongoClient, config.MySQLClient)
	userController := handler.UserAuthController(config.MongoClient, config.MySQLClient)

	// Routes
	e.GET("/feed-user", feedUserController.FeedUser)
	e.POST("/create-feed-user", feedUserController.CreateFeed)

	e.POST("/sign-in", userController.SignIn)
	e.POST("/sign-up", userController.SignUp)

	//Listen
	e.Logger.Fatal(e.Start(":2000"))
}
