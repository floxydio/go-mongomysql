package models

type UserFeed struct {
	NamaUser  string `json:"nama_user" bson:"nama_user" form:"nama_user" query:"nama_user"`
	Feed      string `json:"feed" bson:"feed" form:"feed" query:"feed"`
	Likes     int    `json:"likes" bson:"likes" form:"likes" query:"likes"`
	CreatedAt string `json:"created_at" bson:"created_at" form:"created_at" query:"created_at"`
}
