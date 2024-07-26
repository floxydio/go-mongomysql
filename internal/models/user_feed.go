package models

type UserFeed struct {
	Id        string `json:"_id" bson:"_id"`
	UsersId   int    `json:"users_id" bson:"users_id"`
	Feed      string `json:"feed" bson:"feed" form:"feed" query:"feed"`
	CreatedAt string `json:"created_at" bson:"created_at" form:"created_at" query:"created_at"`
}
