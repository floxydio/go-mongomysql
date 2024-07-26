package models

import "time"

type Users struct {
	UsersId        uint      `gorm:"primaryKey;column:users_id" json:"users_id"`
	Name           string    `json:"name" form:"name"`
	Email          string    `json:"email" form:"email"`
	Password       string    `json:"password" form:"password"`
	DateOfBirth    string    `gorm:"column:date_of_birth" json:"date_of_birth" form:"date_of_birth"`
	ProfilePicture string    `gorm:"profile_picture;default:profile.png" json:"profile_picture" form:"profile_picture"`
	StatusActive   uint      `gorm:"status_active" json:"status_active" form:"status_active"`
	CreatedAt      time.Time `gorm:"column:created_at" json:"created_at" form:"created_at"`
}

type UserLogin struct {
	Email          string `json:"email" form:"email"`
	Password       string `json:"password" form:"password"`
	ProfilePicture string `json:"profile_picture"`
}

type UserData struct {
	UsersId uint   `json:"users_id"`
	Name    string `json:"name"`
}

type ResponseLogin struct {
	StatusCode uint     `json:"status"`
	Data       UserData `json:"data"`
	Token      string   `json:"token"`
	Message    string   `json:"message"`
}

func (Users) TableName() string {
	return "users"
}
