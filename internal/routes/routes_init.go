package routes

import (
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
	"gonosql/internal/config"
	"gonosql/internal/handler"
	"gorm.io/gorm"
)

func RoutesApp(e *echo.Echo, mongoClient *mongo.Database, mysqlClient *gorm.DB) {
	feedUserController := handler.FeedUserController(config.MongoClient, config.MySQLClient)
	userController := handler.UserAuthController(config.MongoClient, config.MySQLClient)
	userLikesController := handler.UserLikesControllerInit(config.MongoClient, config.MySQLClient)

	// === Controller

	e.GET("/feed-user", feedUserController.FeedUser)
	e.POST("/create-feed-user", feedUserController.CreateFeed)

	e.POST("/sign-in", userController.SignIn)
	e.POST("/sign-up", userController.SignUp)

	e.POST("/user-likes", userLikesController.CreateLikes)
}
