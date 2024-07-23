package services

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"gonosql/internal/models"
	"gorm.io/gorm"
)

type FeedUserRepository interface {
	GetFeed() ([]models.UserFeed, error)
	CreateFeed(feed models.UserFeed) error
}

type FeedUserMongoMySQLDB struct {
	mongoClient *mongo.Database
	mySQLClient *gorm.DB
}

func RepositoryFeed(client *mongo.Database, mysqlClients *gorm.DB) FeedUserRepository {
	return &FeedUserMongoMySQLDB{
		mongoClient: client,
		mySQLClient: mysqlClients,
	}
}

func (r FeedUserMongoMySQLDB) GetFeed() ([]models.UserFeed, error) {
	data, err := r.mongoClient.Collection("user_feed").Find(context.Background(), bson.M{})
	fmt.Println("execute getFeed")
	if err != nil {
		return nil, err
	}

	var feeds []models.UserFeed

	if err = data.All(context.Background(), &feeds); err != nil {
		return nil, err
	}
	return feeds, nil
}

func (r FeedUserMongoMySQLDB) CreateFeed(form models.UserFeed) error {
	_, err := r.mongoClient.Collection("user_feed").InsertOne(context.Background(), form)
	if err != nil {
		return err
	}
	return nil
}
