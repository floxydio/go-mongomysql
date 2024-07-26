package handler

import (
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
	"gonosql/internal/models"
	"gonosql/internal/services"
	"gorm.io/gorm"
)

type UserLikesRepository struct {
	likesRepo services.UserLikesInterface
}

func UserLikesControllerInit(mongo *mongo.Database, mysqlClient *gorm.DB) UserLikesRepository {
	return UserLikesRepository{
		likesRepo: services.UserLikesRepository(mongo, mysqlClient),
	}
}

func (userLikesRepo UserLikesRepository) CreateLikes(c echo.Context) error {
	var userLikes models.UserLikes

	if err := c.Bind(&userLikes); err != nil {
		return c.JSON(400, echo.Map{
			"status":  400,
			"error":   true,
			"message": "Failed to bind data",
		})
	}

	errCreate := userLikesRepo.likesRepo.Create(c.Request().Context(), userLikes)

	if errCreate != nil {
		return c.JSON(400, echo.Map{
			"status":  400,
			"error":   true,
			"message": errCreate.Error(),
		})
	}

	return c.JSON(201, echo.Map{
		"status":  201,
		"error":   false,
		"message": "Successfully Create Likes",
	})
}
