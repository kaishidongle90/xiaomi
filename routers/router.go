package routers

import (
	"github.com/astaxie/beego"
	"xiaomi/controllers"
	"xiaomi/controllers/admin"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	var ns = beego.NewNamespace("/admin",
		//beego.NSBefore(func(ctx *context.Context) {
		//	fmt.Println("this is admin")
		//}),
		beego.NSRouter("/",&admin.MainController{}),
		beego.NSRouter("/welcome",&admin.MainController{},"get:Welcome"),
		beego.NSRouter("/login", &admin.LoginController{}),
		beego.NSRouter("/login/doLogin", &admin.LoginController{},"post:DoLogin"),
		beego.NSNamespace("/manager",
			beego.NSInclude(
				&admin.ManagerController{},
			),
		),
		beego.NSNamespace("/role",
			beego.NSInclude(
				&admin.RoleController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
