package controllers

import (
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
)

type BaseController struct {
	web.Controller
}

func (c *BaseController) Prepare() {
	logs.Info("request url :", c.Ctx.Request.URL)
}

func (c *BaseController) Finish() {
	c.ServeJSON(true)
}
