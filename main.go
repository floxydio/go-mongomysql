package main

import (
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	"gonosql/internal/config"
	"gonosql/internal/routes"
	"log"

	_ "github.com/swaggo/echo-swagger/example/docs"
)

// @title Social Media API
// @version 1.0

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost
// @BasePath /
func main() {
	e := echo.New()
	errConnectEnv := godotenv.Load()

	if errConnectEnv != nil {
		log.Fatal("Error loading .env file")
	}
	config.ConnectMongo()
	config.ConnectMysql()

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// Routes
	routes.RoutesApp(e, config.MongoClient, config.MySQLClient)

	//Listen
	e.Logger.Fatal(e.Start(":2000"))
}
