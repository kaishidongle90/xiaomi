package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["xiaomi/controllers/admin:AccessController"] = append(beego.GlobalControllerRouter["xiaomi/controllers/admin:AccessController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["xiaomi/controllers/admin:AccessController"] = append(beego.GlobalControllerRouter["xiaomi/controllers/admin:AccessController"],
        beego.ControllerComments{
            Method: "Add",
            Router: `/add`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["xiaomi/controllers/admin:AccessController"] = append(beego.GlobalControllerRouter["xiaomi/controllers/admin:AccessController"],
        beego.ControllerComments{
            Method: "DoAdd",
            Router: `/add`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["xiaomi/controllers/admin:AccessController"] = append(beego.GlobalControllerRouter["xiaomi/controllers/admin:AccessController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/delete`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["xiaomi/controllers/admin:AccessController"] = append(beego.GlobalControllerRouter["xiaomi/controllers/admin:AccessController"],
        beego.ControllerComments{
            Method: "Edit",
            Router: `/edit`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["xiaomi/controllers/admin:AccessController"] = append(beego.GlobalControllerRouter["xiaomi/controllers/admin:AccessController"],
        beego.ControllerComments{
            Method: "DoEdit",
            Router: `/edit`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["xiaomi/controllers/admin:ManagerController"] = append(beego.GlobalControllerRouter["xiaomi/controllers/admin:ManagerController"],
        beego.ControllerComments{
            Method: "Manager",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["xiaomi/controllers/admin:ManagerController"] = append(beego.GlobalControllerRouter["xiaomi/controllers/admin:ManagerController"],
        beego.ControllerComments{
            Method: "Add",
            Router: `/add`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["xiaomi/controllers/admin:ManagerController"] = append(beego.GlobalControllerRouter["xiaomi/controllers/admin:ManagerController"],
        beego.ControllerComments{
            Method: "DoAdd",
            Router: `/add`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["xiaomi/controllers/admin:ManagerController"] = append(beego.GlobalControllerRouter["xiaomi/controllers/admin:ManagerController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/delete`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["xiaomi/controllers/admin:ManagerController"] = append(beego.GlobalControllerRouter["xiaomi/controllers/admin:ManagerController"],
        beego.ControllerComments{
            Method: "Edit",
            Router: `/edit`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["xiaomi/controllers/admin:ManagerController"] = append(beego.GlobalControllerRouter["xiaomi/controllers/admin:ManagerController"],
        beego.ControllerComments{
            Method: "DoEdit",
            Router: `/edit`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["xiaomi/controllers/admin:RoleController"] = append(beego.GlobalControllerRouter["xiaomi/controllers/admin:RoleController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["xiaomi/controllers/admin:RoleController"] = append(beego.GlobalControllerRouter["xiaomi/controllers/admin:RoleController"],
        beego.ControllerComments{
            Method: "Add",
            Router: `/add`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["xiaomi/controllers/admin:RoleController"] = append(beego.GlobalControllerRouter["xiaomi/controllers/admin:RoleController"],
        beego.ControllerComments{
            Method: "DoAdd",
            Router: `/add`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["xiaomi/controllers/admin:RoleController"] = append(beego.GlobalControllerRouter["xiaomi/controllers/admin:RoleController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/delete`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["xiaomi/controllers/admin:RoleController"] = append(beego.GlobalControllerRouter["xiaomi/controllers/admin:RoleController"],
        beego.ControllerComments{
            Method: "Edit",
            Router: `/edit`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["xiaomi/controllers/admin:RoleController"] = append(beego.GlobalControllerRouter["xiaomi/controllers/admin:RoleController"],
        beego.ControllerComments{
            Method: "DoEdit",
            Router: `/edit`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
