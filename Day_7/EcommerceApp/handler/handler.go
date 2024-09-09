package handler

import (
	"errors"
	"net/http"
	"regexp"
	"time"

	"github.com/PraveenRajPurak/Learning_Golang/Day_7/EcommerceApp/modules/auth"
	"github.com/PraveenRajPurak/Learning_Golang/Day_7/EcommerceApp/modules/config"
	"github.com/PraveenRajPurak/Learning_Golang/Day_7/EcommerceApp/modules/database"
	"github.com/PraveenRajPurak/Learning_Golang/Day_7/EcommerceApp/modules/database/query"
	"github.com/PraveenRajPurak/Learning_Golang/Day_7/EcommerceApp/modules/encrypt"
	"github.com/PraveenRajPurak/Learning_Golang/Day_7/EcommerceApp/modules/model"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type GoApp struct {
	App *config.GoAppTools
	DB  database.DBRepo
}

func NewGoApp(app *config.GoAppTools, db *mongo.Client) *GoApp {
	return &GoApp{
		App: app,
		DB:  query.NewGoAppDB(app, db),
	}
}

func (ga *GoApp) Home() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Welcome to the home page of Ecommerce App!",
		})
	}
}

func (ga *GoApp) Sign_Up() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var user *model.User

		err := ctx.ShouldBindJSON(&user)
		if err != nil {
			_ = ctx.AbortWithError(http.StatusBadRequest, gin.Error{
				Err: err,
			})
		}

		user.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

		user.Password, _ = encrypt.Hash(user.Password)

		if err != nil {
			_ = ctx.AbortWithError(http.StatusInternalServerError, gin.Error{Err: err})
		}

		if err := ga.App.Validate.Struct(&user); err != nil {
			if _, ok := err.(*validator.InvalidValidationError); !ok {
				_ = ctx.AbortWithError(http.StatusBadRequest, gin.Error{Err: err})
				ga.App.InfoLogger.Println(err)
				return
			}
		}

		ok, status, err := ga.DB.InsertUser(user)

		if err != nil {
			_ = ctx.AbortWithError(http.StatusInternalServerError, errors.New("error while adding new user"))
			ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		if !ok {
			_ = ctx.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		switch status {
		case 1:
			{
				ctx.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
			}
		case 2:
			{
				ctx.JSON(http.StatusConflict, gin.H{"message": "User already exists"})
			}
		}
	}
}

func (ga *GoApp) Sign_In() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var user *model.User
		if err := ctx.ShouldBindJSON(&user); err != nil {
			_ = ctx.AbortWithError(http.StatusBadRequest, gin.Error{Err: err})
		}

		regMail := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
		ok := regMail.MatchString(user.Email)

		if ok {

			res, err := ga.DB.VerifyUser(user.Email)
			if err != nil {
				_ = ctx.AbortWithError(http.StatusInternalServerError, err)
				ctx.JSON(http.StatusUnauthorized, gin.H{"message": "unregistered user"})
				return
			}

			id := res["_id"].(primitive.ObjectID)
			password := res["password"].(string)

			verified, err := encrypt.VerifyPassword(user.Password, password)
			if err != nil {
				_ = ctx.AbortWithError(http.StatusInternalServerError, err)
				ctx.JSON(http.StatusUnauthorized, gin.H{"message": "unregistered user detected using wrong password"})
				return
			}

			if verified {

				cookieData := sessions.Default(ctx)

				userInfo := map[string]interface{}{
					"ID":    id,
					"Email": user.Email,
					"Role":  res["role"],
					"Name":  res["name"],
				}

				cookieData.Set("userInfo", userInfo)
				if err := cookieData.Save(); err != nil {
					_ = ctx.AbortWithError(http.StatusInternalServerError, err)
					ctx.JSON(http.StatusInternalServerError, gin.H{"message": "error while saving cookie"})
					return
				}

				t1, t2, err := auth.Generate(user.Email, id, res["role"].(string), res["name"].(string))

				if err != nil {
					_ = ctx.AbortWithError(http.StatusInternalServerError, err)
					ctx.JSON(http.StatusInternalServerError, gin.H{"message": "error while generating tokens"})
					return
				}

				cookieData.Set("token", t1)

				if err := cookieData.Save(); err != nil {
					_ = ctx.AbortWithError(http.StatusInternalServerError, err)
					ctx.JSON(http.StatusInternalServerError, gin.H{"message": "error while saving cookie"})
					return
				}

				cookieData.Set("new_token", t2)

				if err := cookieData.Save(); err != nil {
					_ = ctx.AbortWithError(http.StatusInternalServerError, err)
					ctx.JSON(http.StatusInternalServerError, gin.H{"message": "error while saving cookie"})
					return
				}

				tk := map[string]string{
					"token":    t1,
					"newToken": t2,
				}

				updated, err := ga.DB.UpdateUser(id, tk)

				if err != nil {
					_ = ctx.AbortWithError(http.StatusInternalServerError, err)
					ctx.JSON(http.StatusInternalServerError, gin.H{"message": "error while updating tokens"})
					return
				}

				if !updated {
					_ = ctx.AbortWithError(http.StatusInternalServerError, err)
					ctx.JSON(http.StatusInternalServerError, gin.H{"message": "error while updating tokens"})
					return
				}

				ctx.JSON(http.StatusOK, gin.H{
					"message":       "Successfully Logged in",
					"email":         user.Email,
					"id":            id,
					"role":          res["role"],
					"name":          res["name"],
					"session_token": t1,
				})
			} else {
				ctx.JSON(http.StatusUnauthorized, gin.H{"message": "unregistered user detected using wrong credentials"})
				return
			}
		}
	}
}

func (ga *GoApp) ForgotPassword() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		email, _ := ctx.Get("email")

		var user *model.User

		if err := ctx.ShouldBindJSON(&user); err != nil {
			_ = ctx.AbortWithError(http.StatusBadRequest, gin.Error{Err: err})
		}

		user.Email = email.(string)

		updated, err := ga.DB.CreateNewPassword(user.Email, user.Password)

		if err != nil {
			_ = ctx.AbortWithError(http.StatusInternalServerError, err)
		}

		if !updated {
			_ = ctx.AbortWithError(http.StatusInternalServerError, err)
		}

		ctx.JSON(http.StatusOK, gin.H{"message": "password changed successfully"})
	}
}

func (g *GoApp) InsertProducts() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var product *model.Product

		supplier_id, _ := ctx.Get("ID")

		if err := ctx.ShouldBindJSON(&product); err != nil {
			_ = ctx.AbortWithError(http.StatusBadRequest, gin.Error{Err: err})
		}

		product.SupplierID = supplier_id.(primitive.ObjectID)

		filter := bson.D{{Key: "Name", Value: product.Name}}

	}
}

func (g *GoApp) ViewProducts() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Welcome to the update products page of Ecommerce App!",
		})
	}
}
