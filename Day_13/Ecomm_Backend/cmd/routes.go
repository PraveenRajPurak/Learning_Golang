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

	protected := r.Group("/")
	protected.Use(Authorisation())

	protected.POST("/forgot-password", g.ForgotPassword())
	protected.POST("/insert-products", g.InsertProducts())
	protected.GET("/view-products", g.ViewProducts())

}
