package main

import (
	"MicroNet.Logger/routers"
	"github.com/gin-gonic/gin"
)

func initRouter() *gin.Engine {
	router := gin.Default()

	//用户路由
	routers.LogRouter(router)

	//公共路由
	//routers.CommonRouter(router)

	return router
}
