package query

import (
	"context"
	"errors"
	"fmt"
	"regexp"
	"time"

	"github.com/PraveenRajPurak/Learning_Golang/Day_7/EcommerceApp/modules/config"
	"github.com/PraveenRajPurak/Learning_Golang/Day_7/EcommerceApp/modules/encrypt"
	"github.com/PraveenRajPurak/Learning_Golang/Day_7/EcommerceApp/modules/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type GoAppDB struct {
	App *config.GoAppTools
	DB  *mongo.Client
}

func NewGoAppDB(app *config.GoAppTools, db *mongo.Client) *GoAppDB {
	return &GoAppDB{
		App: app,
		DB:  db,
	}
}

func (g *GoAppDB) InsertUser(user *model.User) (bool, int, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)

	defer cancel()

	regMail := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

	if !regMail.MatchString(user.Email) {

		g.App.ErrorLogger.Println("invalid registered details - email")
		return false, 0, errors.New("invalid registered details - email")

	}

	filter := bson.D{{Key: "email", Value: user.Email}}

	var res bson.M
	err := User(g.DB, "user").FindOne(ctx, filter).Decode(&res)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			user.ID = primitive.NewObjectID()
			_, insertErr := User(g.DB, "user").InsertOne(ctx, user)
			if insertErr != nil {
				g.App.ErrorLogger.Fatalf("cannot add user to the database : %v ", insertErr)
			}
			return true, 1, nil
		}
		g.App.ErrorLogger.Fatal(err)
	}
	return true, 2, nil
}

func (g *GoAppDB) VerifyUser(email string) (primitive.M, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	var res bson.M

	filter := bson.D{{Key: "email", Value: email}}
	err := User(g.DB, "user").FindOne(ctx, filter).Decode(&res)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			g.App.ErrorLogger.Println("no document found for this query")
			return nil, err
		}
		g.App.ErrorLogger.Fatalf("cannot execute the database query perfectly : %v ", err)
	}

	return res, nil
}

func (g *GoAppDB) UpdateUser(userID primitive.ObjectID, tk map[string]string) (bool, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)

	defer cancel()

	filter := bson.D{{Key: "_id", Value: userID}}
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "token", Value: tk["t1"]}, {Key: "new_token", Value: tk["t2"]}}}}

	_, err := User(g.DB, "user").UpdateOne(ctx, filter, update)
	if err != nil {
		g.App.ErrorLogger.Fatalf("cannot update user's tokens in the database : %v ", err)
		return false, err
	}
	return true, nil
}

func (g *GoAppDB) InsertProduct(product *model.Product) (bool, int, error) {

	suppID := product.SupplierID
	fmt.Println("Inserting product...")

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)

	defer cancel()

	filter := bson.D{{Key: "name", Value: product.Name}, {Key: "supplier_id", Value: suppID}}

	var res bson.M

	err := Product(g.DB, "product").FindOne(ctx, filter).Decode(&res)

	if err != nil {

		if err == mongo.ErrNoDocuments {

			product.ID = primitive.NewObjectID()
			_, insertErr := Product(g.DB, "product").InsertOne(ctx, product)
			if insertErr != nil {
				g.App.ErrorLogger.Fatalf("cannot add product to the database : %v ", insertErr)
			}

			return true, 1, nil
		}

		g.App.ErrorLogger.Fatalf("cannot execute the database query perfectly : %v ", err)
	}

	return true, 2, nil

}

func (g *GoAppDB) CreateNewPassword(email string, password string) (bool, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)

	hashed_Password, err := encrypt.Hash(password)

	if err != nil {
		g.App.ErrorLogger.Fatalf("cannot hash password : %v ", err)
		return false, err
	}

	defer cancel()

	filter := bson.D{{Key: "email", Value: email}}
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "password", Value: hashed_Password}}}}

	_, err = User(g.DB, "user").UpdateOne(ctx, filter, update)
	if err != nil {
		g.App.ErrorLogger.Fatalf("cannot update user's password in the database : %v ", err)
		return false, err
	}

	fmt.Println("Creating new password...")
	return true, nil
}

func (g *GoAppDB) ViewProducts() ([]primitive.M, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	var res []primitive.M
	cursor, err := Product(g.DB, "product").Find(ctx, bson.D{})
	if err != nil {
		g.App.ErrorLogger.Fatalf("cannot execute the database query perfectly : %v ", err)
		return nil, err
	}

	if err = cursor.All(ctx, &res); err != nil {
		g.App.ErrorLogger.Fatalf("cannot execute the database query perfectly : %v ", err)
		return nil, err
	}

	return res, nil
}
