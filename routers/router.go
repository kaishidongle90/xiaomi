package routers

import (
	"github.com/astaxie/beego"
	"xiaomi/controllers"
	"xiaomi/controllers/admin"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/qr", &admin.QRCodeController{},"get:QR")
	var ns = beego.NewNamespace("/admin",
		//beego.NSBefore(func(ctx *context.Context) {
		//	fmt.Println("this is admin")
		//}),
		beego.NSRouter("/",&admin.MainController{}),
		beego.NSRouter("/main/changestatus",&admin.MainController{},"get:ChangeStatus"),
		beego.NSRouter("/main/editnum",&admin.MainController{},"get:EditNum"),
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
		beego.NSNamespace("/focus", beego.NSInclude(&admin.FocusController{})),
		beego.NSNamespace("/goodscate", beego.NSInclude(&admin.GoodsCateController{})),
		beego.NSNamespace("/goodstype", beego.NSInclude(&admin.GoodsTypeController{})),
		beego.NSNamespace("/goodstype", beego.NSInclude(&admin.GoodsTypeController{})),
		beego.NSNamespace("/goodstypeAttribute", beego.NSInclude(&admin.GoodsTypeAttrController{})),
		beego.NSNamespace("/goods", beego.NSInclude(&admin.GoodsController{})),
		//beego.NSRouter("/auth", &admin.RoleController{},"get:Auth;post:DoAuth"),
		)
	beego.AddNamespace(ns)
}
