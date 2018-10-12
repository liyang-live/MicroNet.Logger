package routers

import (
	"MicroNet.Logger/controller"
	"github.com/gin-gonic/gin"
)

func LogRouter(router *gin.Engine) {

	//添加日志
	router.POST("/logging", controller.RecordLog)

	//获取日志
	router.GET("/", controller.GetLogRecord)

	//router.GET("/", controllers.Hello)
	//router.GET("/user", controllers.GetUser)
	//
	//router.GET("/user/all", controllers.GetAllUser)

}
