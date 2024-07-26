package handler

import (
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
	"gonosql/internal/models"
	"gonosql/internal/pkg/jwt"
	"gonosql/internal/services"
	"gorm.io/gorm"
)

type RepositoryFeedController struct {
	repo services.FeedUserRepository
}

func FeedUserController(mongoClient *mongo.Database, mysqlClient *gorm.DB) RepositoryFeedController {
	return RepositoryFeedController{
		repo: services.RepositoryFeed(mongoClient, mysqlClient),
	}
}

func (userFeedRepo RepositoryFeedController) FeedUser(c echo.Context) error {
	data, err := userFeedRepo.repo.GetFeed(c.Request().Context())

	if err != nil {
		return c.JSON(400, echo.Map{
			"status":  400,
			"error":   true,
			"message": err.Error(),
		})
	}

	return c.JSON(200, echo.Map{
		"status":  200,
		"error":   false,
		"data":    data,
		"message": "Successfully Get User Feed",
	})
}

func (userFeedRepo RepositoryFeedController) CreateFeed(c echo.Context) error {
	token, errTkn := jwt.GetTokenFromAuthorization(c)

	if errTkn != nil {
		return c.JSON(401, echo.Map{
			"status":  401,
			"error":   true,
			"message": errTkn.Error(),
		})
	}

	dataToken, errClaim := jwt.GetDataToken(token)

	if errClaim != nil {
		return c.JSON(500, echo.Map{
			"status":  500,
			"error":   true,
			"message": errClaim.Error(),
		})
	}

	var feed models.UserFeed

	feed.UsersId = int(dataToken.UserId)

	if err := c.Bind(&feed); err != nil {
		return c.JSON(400, echo.Map{
			"status":  400,
			"error":   true,
			"message": "Failed to bind data",
		})
	}

	err := userFeedRepo.repo.CreateFeed(c.Request().Context(), feed)

	if err != nil {
		return c.JSON(400, echo.Map{
			"status":  400,
			"error":   true,
			"message": err.Error(),
		})
	}

	return c.JSON(200, echo.Map{
		"status":  200,
		"error":   false,
		"message": "Successfully  Create Feed",
	})
}
