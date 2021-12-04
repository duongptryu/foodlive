package main

import (
	"fooddelivery/component"
	"github.com/gin-gonic/gin"
)

func setupRouter(r *gin.Engine, appCtx component.AppContext) {
	v1Route(r, appCtx)
}

func v1Route (r *gin.Engine, appCtx component.AppContext) {

}