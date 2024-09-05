package query

import (
	"github.com/PraveenRajPurak/Learning_Golang/Day_3/BackendwithGo/modules/config"
	"github.com/PraveenRajPurak/Learning_Golang/Day_3/BackendwithGo/modules/database"
	"github.com/PraveenRajPurak/Learning_Golang/Day_3/BackendwithGo/modules/model"
	"go.mongodb.org/mongo-driver/mongo"
)

type GoAppDB struct {
	App *config.GoAppTools
	DB  *mongo.Client
}

func NewGoAppDB(app *config.GoAppTools, db *mongo.Client) database.DBRepo {
	return &GoAppDB{
		App: app,
		DB:  db,
	}
}

func (g *GoAppDB) InsertUser(user *model.User) (bool, int, error) {
	return false, 0, nil
}
