package database

import "github.com/PraveenRajPurak/Learning_Golang/Day_3/BackendwithGo/modules/model"

type DBRepo interface {
	InsertUser(user *model.User) (bool, int, error)
}
