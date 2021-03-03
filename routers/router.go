package routers

import (
	"beego_example/controllers"
	"beego_example/controllers/user"
	"encoding/json"
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
)

func init() {

	// 添加日志拦截器
	var FilterLog = func(ctx *context.Context) {
		url, _ := json.Marshal(ctx.Input.Data()["RouterPattern"])
		params, _ := json.Marshal(ctx.Request.Form)
		outputBytes, _ := json.Marshal(ctx.Input.Data()["json"])
		divider := " - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -"
		topDivider := "┌" + divider
		middleDivider := "├" + divider
		bottomDivider := "└" + divider
		outputStr := "\n" + topDivider + "\n│ 请求地址:" + string(url) + "\n" + middleDivider + "\n│ 请求参数:" + string(params) + "\n│ 返回数据:" + string(outputBytes) + "\n" + bottomDivider
		logs.Info(outputStr)
	}
	web.InsertFilter("/*", web.FinishRouter, FilterLog)

	ns := web.NewNamespace("/api",
		web.NSNamespace("/user",
			web.NSInclude(
				&user.Controller{},
			),
		),
	)
	web.ErrorController(&controllers.ErrorController{})
	web.AddNamespace(ns)
}
