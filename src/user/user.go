package user

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID               primitive.ObjectID `json:"_id" bson:"_id"`
	Name             string             `json:"name" bson:"name"`
	Username         string             `json:"username" bson:"username"`
	Email            string             `json:"email" bson:"email"`
	Address          string             `json:"address" bson:"address"`
	RegistrationDate string             `json:"registration_date" bson:"registration_date"`
}

type UserList struct {
	Users []User `json:"users" bson:"users"`
}
