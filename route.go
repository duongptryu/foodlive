package main

import (
	"fooddelivery/component"
	"fooddelivery/middleware"
	"fooddelivery/modules/user/usertransport/ginuser"
	"github.com/gin-gonic/gin"
)

func setupRouter(r *gin.Engine, appCtx component.AppContext) {
	r.Use(middleware.Recover(appCtx))
	v1Route(r, appCtx)
}

func v1Route(r *gin.Engine, appCtx component.AppContext) {
	v1 := r.Group("/api/v1")
	{
		v1.POST("/register", ginuser.UserReigster(appCtx))
		v1.POST("/activate", ginuser.UserActiveAccount(appCtx))
		v1.POST("/login", ginuser.UserLogin(appCtx))
	}
}
