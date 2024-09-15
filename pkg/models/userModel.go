package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id       primitive.ObjectID `json:_i`
	UserName string             `json:"username"`
	Password string             `json:"password"`
	Email    string             `json:"email"`
}
