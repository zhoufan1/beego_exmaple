package main

import (
	"beego_example/constants"
	_ "beego_example/routers"
	"beego_example/utils/mysql_utils"
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/filter/cors"
	"path/filepath"
)

func main() {

	//初始化日志
	LogInit()

	//初始化数据库
	DbInit()

	//关闭长连接
	web.BeeApp.Server.SetKeepAlivesEnabled(false)

	logs.Info("project  name start ... ", web.BConfig.AppName)
	logs.Info("project  profile : ", web.BConfig.RunMode)
	if web.BConfig.RunMode == "dev" {
		web.BConfig.WebConfig.DirectoryIndex = true
		web.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	// 允许跨域
	web.InsertFilter("*", web.BeforeRouter, cors.Allow(&cors.Options{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"*"},
		AllowCredentials: true,
	}))

	web.Run()
}

func DbInit() {
	logs.Info("初始化 mysql ")
	mysql_utils.Register(constants.DbUrl, constants.DbMaxIdle, constants.DbMaxOpen, "default")
	//开启自动创表
	orm.RunSyncdb("default", false, true)
	orm.Debug = true
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
