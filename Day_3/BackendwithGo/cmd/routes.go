package main

import (
	"github.com/PraveenRajPurak/Learning_Golang/Day_3/BackendwithGo/handlers"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine, g *handlers.GoApp) {

	router := r.Use(gin.Logger(), gin.Recovery())

	cookieData := cookie.NewStore([]byte("go-app"))
	router.Use(sessions.Sessions("go-app", cookieData))

	router.GET("/", g.Home())

	router.POST("/signup", g.Sign_Up())

	router.POST("/signin", g.SignIn())
}
