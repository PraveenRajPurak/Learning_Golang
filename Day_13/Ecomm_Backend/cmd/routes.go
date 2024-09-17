package main

import (
	"github.com/PraveenRajPurak/Learning_Golang/Day_13/CarsGo/handler"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine, g *handler.GoApp) {
	router := r.Use(gin.Logger(), gin.Recovery())

	cookieData := cookie.NewStore([]byte("ecomm"))
	router.Use(sessions.Sessions("ecomm", cookieData))

	router.GET("/", g.Home())

	router.POST("/sign-up", g.Sign_Up())
	router.POST("/sign-in", g.Sign_In())

	router.POST("/sign-up-admin", g.Sign_Up_Admin())
	router.POST("/sign-in-admin", g.Sign_In_Admin())

	protectedUsers := r.Group("/users")
	protectedUsers.Use(Authorisation())

	protectedUsers.POST("/forgot-password", g.ForgotPasswordUser())
	protectedUsers.GET("/view-products", g.ViewProducts())
	protectedUsers.POST("update-email", g.Update_Email_User())
	protectedUsers.POST("update-name", g.Update_Name_User())
	protectedUsers.POST("update-phone", g.Update_Phone_User())

	protectedAdmin := r.Group("/admin")
	protectedAdmin.Use(Admin_Authorisation())
	protectedAdmin.POST("forgot-password", g.ForgotPasswordAdmin())
	protectedAdmin.POST("create-category", g.CreateCategory())
	protectedAdmin.POST("create-product", g.InsertProducts())
	protectedAdmin.POST("update-email", g.Update_Email_Admin())
	protectedAdmin.POST("update-name", g.Update_Name_Admin())
	protectedAdmin.POST("update-phone", g.Update_Phone_Admin())

}
