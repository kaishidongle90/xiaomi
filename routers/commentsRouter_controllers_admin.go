package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["xiaomi/controllers/admin:AdminController"] = append(beego.GlobalControllerRouter["xiaomi/controllers/admin:AdminController"],
        beego.ControllerComments{
            Method: "Manager",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["xiaomi/controllers/admin:AdminController"] = append(beego.GlobalControllerRouter["xiaomi/controllers/admin:AdminController"],
        beego.ControllerComments{
            Method: "Add",
            Router: `/add`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["xiaomi/controllers/admin:AdminController"] = append(beego.GlobalControllerRouter["xiaomi/controllers/admin:AdminController"],
        beego.ControllerComments{
            Method: "Edit",
            Router: `/edit`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
