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
		beego.NSRouter("/exit", &admin.LoginController{},"get:LoginOut"),
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
		beego.NSNamespace("/access",
			beego.NSInclude(
				&admin.AccessController{},
			),
		),
		//beego.NSRouter("/auth", &admin.RoleController{},"get:Auth;post:DoAuth"),
	)
	beego.AddNamespace(ns)
}
