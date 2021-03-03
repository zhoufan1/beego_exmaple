package constants

import (
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
	"strings"
)

var (
	AppName = web.AppConfig.DefaultString("APPNAME", "beego_example")

	// 数据库创建
	DbMaxIdle = web.AppConfig.DefaultInt("DbMaxIdle", 10)
	DbMaxOpen = web.AppConfig.DefaultInt("DbMaxOpen", 30)
	DbUrl     = web.AppConfig.DefaultString("DbUrl", "")

	//日志
	ExampleLogDir        = web.AppConfig.DefaultString("EXAMPLE_LOG_DIR", "logs")
	ExampleLogLevel      = web.AppConfig.DefaultString("EXAMPLE_LOG_LEVEL", "debug")
	ExampleLogMaxsize    = web.AppConfig.DefaultInt("EXAMPLE_LOG_MAXSIZE", 500)
	ExampleLogMaxbackups = web.AppConfig.DefaultInt("EXAMPLE_LOG_MAXBACKUPS", 0)
	ExampleLogMaxage     = web.AppConfig.DefaultInt("EXAMPLE_LOG_MAXAGE", 15)
)

// SetLevel sets the global log level used by the simple logger.
func GetLevelStr(l string) int {
	switch strings.ToLower(l) {
	case "emergency":
		return logs.LevelEmergency
	case "alert":
		return logs.LevelAlert
	case "critical":
		return logs.LevelCritical
	case "error":
		return logs.LevelError
	case "warn":
		return logs.LevelWarning
	case "notice":
		return logs.LevelNotice
	case "info":
		return logs.LevelInformational
	default:
		return logs.LevelDebug
	}
}
