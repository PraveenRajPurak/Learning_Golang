package query

import "go.mongodb.org/mongo-driver/mongo"

func User(db *mongo.Client, collection string) *mongo.Collection {
	var user = db.Database("Go_App").Collection(collection)
	return user
}
