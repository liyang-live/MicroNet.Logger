// LogController
package controller

import (
	"net/http"
	"time"

	"MicroNet.Logger/common"
	"MicroNet.Logger/models"
	"github.com/gin-gonic/gin"
)

func GetLogRecord(c *gin.Context) {
	c.JSON(http.StatusOK,
		models.LogModel{Id: "1", System: "测试系统", Module: "测试模块", FuncName: "测试方法",
			FuncParameter: "测试参数", LogTime: common.TimFormat(time.Now().String(), "yyyy-MM-dd HH:mm:ss"),
			LogContent: "测试内容"})
}

func RecordLog(c *gin.Context) {

	var logger models.LogModel
	c.ShouldBind(&logger)

	go writeLog(logger)
}

func writeLog(model models.LogModel) {
	//fmt.Println(model)
	//log.Print(model)
	common.WriterFile(model)
}
