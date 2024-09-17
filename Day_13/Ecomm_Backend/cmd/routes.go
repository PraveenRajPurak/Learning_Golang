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

	protectedUsers := r.Group("/users")
	protectedUsers.Use(Authorisation())

	protectedUsers.POST("/forgot-password", g.ForgotPassword())
	protectedUsers.GET("/view-products", g.ViewProducts())

	protectedAdmin := r.Group("/admin")
	protectedAdmin.Use(Admin_Authorisation())

	protectedAdmin.POST("/sign-up-admin", g.Sign_Up_Admin())
	protectedAdmin.POST("/sign-in-admin", g.Sign_In_Admin())
	protectedAdmin.POST("create-category", g.CreateCategory())

}
