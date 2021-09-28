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

// USER-API
type Controller struct {
	controllers.BaseController
}

// @Title Login
// @Description 登录测试
// @Param	userName query string true "用户名"
// @Param	userName query string true "用户密码"
// @Success 200 {object} base.Model
// @router /login [get]
func (u *Controller) Login() {
	userName := u.GetString("userName")
	userPass := u.GetString("userPass")
	user := new(model.User)
	user.UserName = userName
	user.UserPass = userPass
	u.Data["json"] = base.ResponseSuccess(user)
}

// @Title Save
// @Description 保存
// @Param	body body model.User true "保存参数"
// @Success 200 {object} base.Model
// @router /save [post]
func (u *Controller) Save() {
	var user model.User
	if err := json.Unmarshal(u.Ctx.Input.RequestBody, &user); err == nil {
		id, err := service.SaveOrUpdate(&user)
		if err != nil {
			u.Data["json"] = base.ResponseFailure(-1, "添加失败")
			return
		}
		u.Data["json"] = base.ResponseSuccess(id)
	} else {
		logs.Error(err.Error())
		u.Data["json"] = base.ResponseFailure(-1, "参数缺失")
	}
}

// @Title QueryAll
// @Description 根据条件分页查询
// @Param	body body request.UserPagerRequest true "查询条件"
// @Success 200 {object} base.Model
// @router /queryAll [post]
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

// @Title QueryById
// @Description 根据id查询
// @Param	id query true int "主键"
// @Success 200 {object} base.Model
// @router /queryById [get]
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
