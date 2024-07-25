package services

import (
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
	"gonosql/internal/models"
	"gorm.io/gorm"
)

type UsersAuthInterface interface {
	SignUp(form models.Users) error
	SignIn(form models.UserLogin) (models.ResponseLogin, error)
}

type UsersAuthInit struct {
	mongoClient *mongo.Database
	mySQLClient *gorm.DB
}

func UserAuthMongoMySQLDB(clientMongo *mongo.Database, clientMySQL *gorm.DB) UsersAuthInterface {
	return &UsersAuthInit{
		mongoClient: clientMongo,
		mySQLClient: clientMySQL,
	}
}

func (authService UsersAuthInit) SignUp(form models.Users) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(form.Password), 10)
	if err != nil {
		return err
	}
	form.Password = string(hash)
	errRegister := authService.mySQLClient.Create(&form).Error
	if errRegister != nil {
		return errRegister
	}

	return nil
}

func (authService UsersAuthInit) SignIn(form models.UserLogin) (models.ResponseLogin, error) {
	var user models.Users

	err := authService.mySQLClient.Where("email = ?", form.Email).Find(&user).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return models.ResponseLogin{}, fmt.Errorf("user not found")
	}

	if err != nil {
		return models.ResponseLogin{}, err
	}

	errBcrypt := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(form.Password))

	if errBcrypt != nil {
		return models.ResponseLogin{}, fmt.Errorf("%s", "Wrong Username or Password")
	}

	return models.ResponseLogin{
		StatusCode: 200,
		Data: models.UserData{
			UsersId: user.UsersId,
			Name:    user.Name,
		},
		Message: "Successfully Login",
	}, nil

}
