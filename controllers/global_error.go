package controllers

import (
	"beego_example/models/base"
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
)

type ErrorController struct {
	web.Controller
}

func (e *ErrorController) Prepare() {
	logs.Info("error url :", e.Ctx.Request.URL)
}

func (e *ErrorController) Finish() {
	path := e.ViewPath
	logs.Info("path :", path)
	if len(path) <= 0 {
		e.Ctx.ResponseWriter.Status = 200
		e.ServeJSON(true)
	}

}
func (e *ErrorController) Error401() {
	e.Data["json"] = base.ResponseFailure(401, "未认证，请登录")
}

func (e *ErrorController) Error403() {
	e.Data["json"] = base.ResponseFailure(403, "服务不可用")
}

func (e *ErrorController) Error404() {
	e.Data["json"] = base.ResponseFailure(404, "未找到控制器")
}

func (e *ErrorController) Error500() {
	e.Data["json"] = base.ResponseFailure(500, "系统错误")
}

func (e *ErrorController) Error503() {
	e.Data["json"] = base.ResponseFailure(503, "系统错误")
}
