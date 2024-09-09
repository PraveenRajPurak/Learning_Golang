package database

import (
	"github.com/PraveenRajPurak/Learning_Golang/Day_7/EcommerceApp/modules/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DBRepo interface {
	InsertUser(user *model.User) (bool, int, error)
	VerifyUser(email string) (primitive.M, error)
	UpdateUser(userID primitive.ObjectID, tk map[string]string) (bool, error)
	CreateNewPassword(email string, password string) (bool, error)
	InsertPrduct(product *model.Product) (bool, error)
}
