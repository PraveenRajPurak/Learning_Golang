package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `json:"_id" bson:"_id"`
	Name      string             `json:"name" Usage:"required"`
	Email     string             `json:"email" Usage:"required"`
	Password  string             `json:"password" Usage:"required"`
	Role      string             `json:"role"`
	Token     string             `json:"token"`
	New_Token string             `json:"new_token"`
	CreatedAt time.Time          `json:"created_At"`
	UpdatedAt time.Time          `json:"updated_At"`
}

type Product struct {
	ID           primitive.ObjectID `json:"_id" bson:"_id"`
	SupplierID   primitive.ObjectID `json:"supplier_id"`
	Name         string             `json:"name" Usage:"required"`
	Description  string             `json:"description" Usage:"required"`
	Category     string             `json:"category" Usage:"required"`
	RegularPrice int                `json:"regular_price" Usage:"required"`
	SalePrice    int                `json:"sale_price"`
	SaleStarts   time.Time          `json:"sale_starts"`
	SaleEnds     time.Time          `json:"sale_ends"`
	Stock        int                `json:"quantity" Usage:"required"`
	CreatedAt    time.Time          `json:"created_At"`
	UpdatedAt    time.Time          `json:"updated_At"`
}
