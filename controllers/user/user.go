package user

import (
	"beego_example/controllers"
	"beego_example/models"
)

type Controller struct {
	controllers.BaseController
}

type User struct {
	UserName string `json:"userName"`
	UserPass string `json:"userPass"`
}

// @router /login  [get]
func (u *Controller) Login() {
	userName := u.GetString("userName")
	userPass := u.GetString("userPass")
	user := new(User)
	user.UserName = userName
	user.UserPass = userPass
	u.Data["json"] = models.ResponseSuccess(user)
	u.ServeJSON()
}
