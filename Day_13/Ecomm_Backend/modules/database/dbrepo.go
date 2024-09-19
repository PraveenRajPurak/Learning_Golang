package database

import (
	"github.com/PraveenRajPurak/Learning_Golang/Day_13/CarsGo/modules/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DBRepo interface {
	InsertUser(user *model.User) (bool, int, error)
	VerifyUser(email string) (primitive.M, error)
	UpdateUser(userID primitive.ObjectID, tk map[string]string) (bool, error)
	CreateNewPassword(email string, password string) (bool, error)
	InsertProduct(product *model.Product) (bool, int, error)
	ViewProducts() ([]primitive.M, error)
	CreateCategory(category *model.Category) (bool, int, error)
	SignUpAdmin(admin *model.Admin) (bool, int, error)
	VerifyAdmin(email string) (primitive.M, error)
	UpdateAdmin(userID primitive.ObjectID, tk map[string]string) (bool, error)
	SignOutAdmin(adminID primitive.ObjectID) (bool, error)
	SignOutUser(userID primitive.ObjectID) (bool, error)
	CreateNewPasswordAdmin(email string, password string) (bool, error)
	UpdateEmailUser(current_email string, new_email string) (bool, error)
	UpdateEmailAdmin(current_email string, new_email string) (bool, error)
	UpdateNameUser(email string, new_name string) (bool, error)
	UpdateNameAdmin(email string, new_name string) (bool, error)
	UpdatePhoneUser(email string, new_phone string) (bool, error)
	UpdatePhoneAdmin(email string, new_phone string) (bool, error)
	UpdateProduct(product *model.Product) (bool, error)
	Toggle_Stock(productID primitive.ObjectID) (bool, error)
	AddProductToWishlist(Product_Id primitive.ObjectID, User_Id primitive.ObjectID) (bool, error)
	RemoveProductFromWishlist(Product_Id primitive.ObjectID, User_Id primitive.ObjectID) (bool, error)
	GetSingleProduct(Id primitive.ObjectID) (primitive.M, error)
	AddToCart(userID primitive.ObjectID, cartItems *model.CartItems) (bool, error)
	RemoveFromCart(userID primitive.ObjectID, productID primitive.ObjectID) (bool, error)
	GetAllUsers() ([]primitive.M, error)
	InitializeUser(userId primitive.ObjectID) (bool, error)
	CreateOrder(order *model.Order) (bool, error)
	FindUserIDWithName(name string) (primitive.ObjectID, error)
	GetAllOrders() ([]primitive.M, error)
	DeleteProduct(id primitive.ObjectID) (bool, error)
}
