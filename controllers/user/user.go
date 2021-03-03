package user

import (
	"beego_example/controllers"
	"beego_example/models/base"
	"beego_example/models/model"
	"beego_example/models/request"
	"beego_example/service"
	"encoding/json"
	"github.com/beego/beego/v2/core/logs"
)

type Controller struct {
	controllers.BaseController
}

// @router /login  [get]
func (u *Controller) Login() {
	userName := u.GetString("userName")
	userPass := u.GetString("userPass")
	user := new(model.User)
	user.UserName = userName
	user.UserPass = userPass
	u.Data["json"] = base.ResponseSuccess(user)
}

// @router /save  [post]
func (u *Controller) Save() {
	var user model.User
	if err := json.Unmarshal(u.Ctx.Input.RequestBody, &user); err == nil {
		err := service.SaveOrUpdate(&user)
		if err != nil {
			u.Data["json"] = base.ResponseFailure(-1, "添加失败")
			return
		}
		u.Data["json"] = base.ResponseSuccess(nil)
	} else {
		logs.Error(err.Error())
		u.Data["json"] = base.ResponseFailure(-1, "参数缺失")
	}
}

// @router /queryAll  [post]
func (u *Controller) QueryAll() {
	var userPagerRequest request.UserPagerRequest
	if err := json.Unmarshal(u.Ctx.Input.RequestBody, &userPagerRequest); err == nil {
		pager := base.Pager{
			PageNo:   userPagerRequest.PageNo,
			PageSize: userPagerRequest.PageSize,
		}
		queryUser := new(model.User)
		queryUser.UserName = userPagerRequest.UserName
		user, err := service.QueryAll(pager, queryUser)
		if err != nil {
			u.Data["json"] = base.ResponseFailure(-1, err.Error())
			return
		}
		u.Data["json"] = base.ResponseSuccess(user)
	} else {
		logs.Error("QueryAll error ", err.Error())
		u.Data["json"] = base.ResponseFailure(-1, err.Error())
	}
}

// @router /queryById  [get]
func (u *Controller) QueryById() {
	id, err := u.GetInt("id")
	if err != nil || id <= 0 {
		u.Data["json"] = base.ResponseFailure(-1, "参数缺失")
		return
	}
	user, err := service.QueryById(id)
	if err != nil {
		u.Data["json"] = base.ResponseFailure(-1, err.Error())
		return
	}
	u.Data["json"] = base.ResponseSuccess(user)
}
