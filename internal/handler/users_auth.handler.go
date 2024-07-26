package handler

import (
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
	"gonosql/internal/models"
	"gonosql/internal/services"
	"gorm.io/gorm"
)

type UserAuthRepo struct {
	userAuthRepo services.UsersAuthInterface
}

func UserAuthController(mongoClient *mongo.Database, mysqlClient *gorm.DB) UserAuthRepo {
	return UserAuthRepo{
		userAuthRepo: services.UserAuthMongoMySQLDB(mongoClient, mysqlClient),
	}
}

func (authService UserAuthRepo) SignUp(c echo.Context) error {
	var users models.Users
	if err := c.Bind(&users); err != nil {
		return c.JSON(400, echo.Map{
			"status":  400,
			"error":   true,
			"message": "Failed to bind data",
		})
	}
	err := authService.userAuthRepo.SignUp(users)

	if err != nil {
		return c.JSON(400, echo.Map{
			"status":  400,
			"error":   true,
			"message": err.Error(),
		})
	}

	return c.JSON(201, echo.Map{
		"status":  201,
		"error":   false,
		"message": "Successfully Create User",
	})
}

func (authService UserAuthRepo) SignIn(c echo.Context) error {
	var userLogin models.UserLogin
	if err := c.Bind(&userLogin); err != nil {
		return c.JSON(400, echo.Map{
			"status":  400,
			"error":   true,
			"message": "Failed to bind data",
		})
	}
	data, err := authService.userAuthRepo.SignIn(userLogin)

	if err != nil {
		return c.JSON(400, echo.Map{
			"status":  400,
			"error":   true,
			"message": err.Error(),
		})
	}

	return c.JSON(200, data)

}
