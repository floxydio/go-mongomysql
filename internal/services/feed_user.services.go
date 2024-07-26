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

type FeedUserRepository interface {
	GetFeed(context context.Context) ([]models.UserFeed, error)
	CreateFeed(context context.Context, feed models.UserFeed) error
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

func (r FeedUserMongoMySQLDB) GetFeed(ctx context.Context) ([]models.UserFeed, error) {
	data, err := r.mongoClient.Collection("user_feed").Find(ctx, bson.M{})
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

func (r FeedUserMongoMySQLDB) CreateFeed(ctx context.Context, form models.UserFeed) error {
	if form.Id == "" {
		form.Id = primitive.NewObjectID().Hex()
	}
	_, err := r.mongoClient.Collection("user_feed").InsertOne(ctx, form)
	if err != nil {
		return err
	}
	return nil
}
