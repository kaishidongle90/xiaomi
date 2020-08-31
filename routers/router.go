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
		beego.NSRouter("/login", &admin.AdminController{}),
		beego.NSNamespace("/manager",
			beego.NSInclude(
				&admin.AdminController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
