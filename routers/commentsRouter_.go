package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {

	beego.GlobalControllerRouter["beego_example/controllers/user:Controller"] = append(beego.GlobalControllerRouter["beego_example/controllers/user:Controller"],
		beego.ControllerComments{
			Method:           "Login",
			Router:           "/login",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["beego_example/controllers/user:Controller"] = append(beego.GlobalControllerRouter["beego_example/controllers/user:Controller"],
		beego.ControllerComments{
			Method:           "QueryAll",
			Router:           "/queryAll",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["beego_example/controllers/user:Controller"] = append(beego.GlobalControllerRouter["beego_example/controllers/user:Controller"],
		beego.ControllerComments{
			Method:           "QueryById",
			Router:           "/queryById",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["beego_example/controllers/user:Controller"] = append(beego.GlobalControllerRouter["beego_example/controllers/user:Controller"],
		beego.ControllerComments{
			Method:           "Save",
			Router:           "/save",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

}
