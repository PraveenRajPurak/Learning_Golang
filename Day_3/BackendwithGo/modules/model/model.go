package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id        primitive.ObjectID `json:"_id" bson:"_id"`
	Name      string             `json:"name" Usage:"required"`
	Email     string             `json:"email" Usage:"required"`
	Password  string             `json:"password" Usage:"required"`
	Role      string             `json:"role"`
	Token     string             `json:"token"`
	CreatedAt time.Time          `json:"created_At"`
	UpdatedAt time.Time          `json:"updated_At"`
}
