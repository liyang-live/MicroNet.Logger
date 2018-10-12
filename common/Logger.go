// Logger
package common

import (
	"MicroNet.Logger/models"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func WriterFile(model models.LogModel) {

	path := filepath.Dir(os.Args[0]) + "/log/"

	_, err := os.Stat(path)
	if err == nil {
		os.Mkdir(path, os.ModePerm)
	}
	if os.IsNotExist(err) {
		os.Mkdir(path, os.ModePerm)
	}

	path = path + TimFormat(time.Now().String(), "yyyy-MM-dd") + ".txt"

	//创建输出日志文件
	logFile, err := os.OpenFile(path, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	defer logFile.Close()
	if err != nil {
		fmt.Println(err)
	}
	log, _ := json.Marshal(model)
	logFile.Write(log)
	logFile.WriteString("\r\n\r\n\r\n")
}
