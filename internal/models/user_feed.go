package models

type UserFeed struct {
	Id        string `json:"_id" bson:"_id"`
	UsersId   int    `json:"users_id"`
	NamaUser  string `json:"nama_user" bson:"nama_user" form:"nama_user" query:"nama_user"`
	Feed      string `json:"feed" bson:"feed" form:"feed" query:"feed"`
	CreatedAt string `json:"created_at" bson:"created_at" form:"created_at" query:"created_at"`
}
