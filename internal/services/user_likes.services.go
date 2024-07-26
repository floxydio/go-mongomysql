package services

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"gonosql/internal/models"
	"gorm.io/gorm"
)

type UserLikesInterface interface {
	FindByUserIdAndFeed(context context.Context, userId uint, feedId string) (int, error)
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

func (userLikes UserLikesService) FindByUserIdAndFeed(context context.Context, userId uint, feedId string) (int, error) {
	var feedModel models.UserFeed
	err := userLikes.mongoClient.Collection("user_likes").FindOne(context, bson.M{
		"_id":      feedId,
		"users_id": userId,
	}).Decode(&feedModel)

	if err != nil {
		return 0, fmt.Errorf("something went wrong when fetching %s", err.Error())
	}

	return 1, nil
}

func (userLikes UserLikesService) Create(context context.Context, form models.UserLikes) error {
	if form.Id == "" {
		form.Id = primitive.NewObjectID().Hex()
	}
	_, err := userLikes.mongoClient.Collection("user_likes").InsertOne(context, form)

	if err != nil {
		return err
	}
	return nil
}
