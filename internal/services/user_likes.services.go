package services

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"gonosql/internal/models"
	"gorm.io/gorm"
)

type UserLikesInterface interface {
	Create(context context.Context, form models.UserLikes) error
}

type UserLikesService struct {
	mongoClient *mongo.Database
	mysqlClient *gorm.DB
}

func UserLikesRepository(mongoClients *mongo.Database, dbClient *gorm.DB) UserLikesInterface {
	return &UserLikesService{
		mongoClient: mongoClients,
		mysqlClient: dbClient,
	}
}

func (userLikes UserLikesService) Create(context context.Context, form models.UserLikes) error {

	_, err := userLikes.mongoClient.Collection("user_likes").InsertOne(context, form)

	if err != nil {
		return err
	}
	return nil
}
