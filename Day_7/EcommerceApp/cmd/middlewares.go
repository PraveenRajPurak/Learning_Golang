package main

import (
	"errors"
	"net/http"

	"github.com/PraveenRajPurak/Learning_Golang/Day_7/EcommerceApp/modules/auth"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Authorisation() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		cookieData := sessions.Default(ctx)

		accessToken := cookieData.Get("token").(string)

		if accessToken == "" {

			_ = ctx.AbortWithError(http.StatusUnauthorized, errors.New("unauthorized user"))
			return
		}

		claims, err := auth.Parse(accessToken)

		if err != nil {
			_ = ctx.AbortWithError(http.StatusUnauthorized, gin.Error{
				Err: err})
		}

		ctx.Set("pass", accessToken)
		ctx.Set("Email", claims.Email)
		ctx.Set("ID", claims.ID)
		ctx.Set("Name", claims.Name)
		ctx.Set("Role", claims.Role)
		ctx.Next()
	}
}
