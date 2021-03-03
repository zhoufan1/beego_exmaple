package main

import (
	"beego_example/constants"
	_ "beego_example/routers"
	"beego_example/utils/mysql_utils"
	"fmt"
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
	"path/filepath"
)

func main() {

	//初始化日志
	LogInit()

	//初始化数据库
	DbInit()

}

func DbInit() {
	logs.Info("project  name start ... ", web.BConfig.AppName)
	logs.Info("project  profile : ", web.BConfig.RunMode)
	web.BConfig.WebConfig.DirectoryIndex = true
	web.BConfig.WebConfig.StaticDir["swagger"] = "/swagger"
	web.Run()
	logs.Info("初始化 mysql ")
	mysql_utils.Register(constants.DbUrl, constants.DbMaxIdle, constants.DbMaxOpen, "default")
}

func LogInit() {
	logs.SetLogFuncCall(true)
	logs.Async()
	filePath := filepath.Join(constants.ExampleLogDir, "service.log")
	logs.Info("filePath :", filePath)
	config := fmt.Sprintf(`{"filename":"%s","separate":["emergency", "alert", "critical", "error", "warning", "notice", "info", "debug"]}`, filePath)
	err := logs.SetLogger(logs.AdapterMultiFile, config)
	logs.Error("logger :", err)
	logs.SetLevel(constants.GetLevelStr(constants.ExampleLogLevel))

}
