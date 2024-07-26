package models

type UserLikes struct {
	Id      string `json:"_id" bson:"_id"`
	FeedId  string `json:"feed_id" bson:"feed_id" form:"feed_id"`
	Likes   uint   `json:"likes" bson:"likes" form:"likes"`
	UsersId uint   `json:"users_id" bson:"users_id" form:"users_id"`
}
