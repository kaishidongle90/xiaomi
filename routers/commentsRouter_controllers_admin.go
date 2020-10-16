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

    beego.GlobalControllerRouter["xiaomi/controllers/admin:FocusController"] = append(beego.GlobalControllerRouter["xiaomi/controllers/admin:FocusController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["xiaomi/controllers/admin:FocusController"] = append(beego.GlobalControllerRouter["xiaomi/controllers/admin:FocusController"],
        beego.ControllerComments{
            Method: "Add",
            Router: `/add`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["xiaomi/controllers/admin:FocusController"] = append(beego.GlobalControllerRouter["xiaomi/controllers/admin:FocusController"],
        beego.ControllerComments{
            Method: "DoAdd",
            Router: `/add`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["xiaomi/controllers/admin:FocusController"] = append(beego.GlobalControllerRouter["xiaomi/controllers/admin:FocusController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/delete`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["xiaomi/controllers/admin:FocusController"] = append(beego.GlobalControllerRouter["xiaomi/controllers/admin:FocusController"],
        beego.ControllerComments{
            Method: "Edit",
            Router: `/edit`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["xiaomi/controllers/admin:FocusController"] = append(beego.GlobalControllerRouter["xiaomi/controllers/admin:FocusController"],
        beego.ControllerComments{
            Method: "DOEdit",
            Router: `/edit`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["xiaomi/controllers/admin:GoodsCateController"] = append(beego.GlobalControllerRouter["xiaomi/controllers/admin:GoodsCateController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["xiaomi/controllers/admin:GoodsCateController"] = append(beego.GlobalControllerRouter["xiaomi/controllers/admin:GoodsCateController"],
        beego.ControllerComments{
            Method: "Add",
            Router: `/add`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["xiaomi/controllers/admin:GoodsCateController"] = append(beego.GlobalControllerRouter["xiaomi/controllers/admin:GoodsCateController"],
        beego.ControllerComments{
            Method: "DoAdd",
            Router: `/add`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["xiaomi/controllers/admin:GoodsCateController"] = append(beego.GlobalControllerRouter["xiaomi/controllers/admin:GoodsCateController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/delete`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["xiaomi/controllers/admin:GoodsCateController"] = append(beego.GlobalControllerRouter["xiaomi/controllers/admin:GoodsCateController"],
        beego.ControllerComments{
            Method: "Edit",
            Router: `/edit`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["xiaomi/controllers/admin:GoodsCateController"] = append(beego.GlobalControllerRouter["xiaomi/controllers/admin:GoodsCateController"],
        beego.ControllerComments{
            Method: "DoEdit",
            Router: `/edit`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["xiaomi/controllers/admin:GoodsController"] = append(beego.GlobalControllerRouter["xiaomi/controllers/admin:GoodsController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["xiaomi/controllers/admin:GoodsController"] = append(beego.GlobalControllerRouter["xiaomi/controllers/admin:GoodsController"],
        beego.ControllerComments{
            Method: "Add",
            Router: `/add`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["xiaomi/controllers/admin:GoodsController"] = append(beego.GlobalControllerRouter["xiaomi/controllers/admin:GoodsController"],
        beego.ControllerComments{
            Method: "DoAdd",
            Router: `/add`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["xiaomi/controllers/admin:GoodsController"] = append(beego.GlobalControllerRouter["xiaomi/controllers/admin:GoodsController"],
        beego.ControllerComments{
            Method: "ChangeGoodsImageColor",
            Router: `/changeGoodsImageColor`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["xiaomi/controllers/admin:GoodsController"] = append(beego.GlobalControllerRouter["xiaomi/controllers/admin:GoodsController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/delete`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["xiaomi/controllers/admin:GoodsController"] = append(beego.GlobalControllerRouter["xiaomi/controllers/admin:GoodsController"],
        beego.ControllerComments{
            Method: "Edit",
            Router: `/edit`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["xiaomi/controllers/admin:GoodsController"] = append(beego.GlobalControllerRouter["xiaomi/controllers/admin:GoodsController"],
        beego.ControllerComments{
            Method: "DoEdit",
            Router: `/edit`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["xiaomi/controllers/admin:GoodsController"] = append(beego.GlobalControllerRouter["xiaomi/controllers/admin:GoodsController"],
        beego.ControllerComments{
            Method: "GetGoodsTypeAttribute",
            Router: `/getGoodsTypeAttribute`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["xiaomi/controllers/admin:GoodsController"] = append(beego.GlobalControllerRouter["xiaomi/controllers/admin:GoodsController"],
        beego.ControllerComments{
            Method: "RemoveGoodsImage",
            Router: `/removeGoodsImage`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["xiaomi/controllers/admin:GoodsController"] = append(beego.GlobalControllerRouter["xiaomi/controllers/admin:GoodsController"],
        beego.ControllerComments{
            Method: "Uplod",
            Router: `/upload`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["xiaomi/controllers/admin:GoodsTypeAttrController"] = append(beego.GlobalControllerRouter["xiaomi/controllers/admin:GoodsTypeAttrController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["xiaomi/controllers/admin:GoodsTypeAttrController"] = append(beego.GlobalControllerRouter["xiaomi/controllers/admin:GoodsTypeAttrController"],
        beego.ControllerComments{
            Method: "Add",
            Router: `/add`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["xiaomi/controllers/admin:GoodsTypeAttrController"] = append(beego.GlobalControllerRouter["xiaomi/controllers/admin:GoodsTypeAttrController"],
        beego.ControllerComments{
            Method: "DoAdd",
            Router: `/add`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["xiaomi/controllers/admin:GoodsTypeAttrController"] = append(beego.GlobalControllerRouter["xiaomi/controllers/admin:GoodsTypeAttrController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/delete`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["xiaomi/controllers/admin:GoodsTypeAttrController"] = append(beego.GlobalControllerRouter["xiaomi/controllers/admin:GoodsTypeAttrController"],
        beego.ControllerComments{
            Method: "Edit",
            Router: `/edit`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["xiaomi/controllers/admin:GoodsTypeAttrController"] = append(beego.GlobalControllerRouter["xiaomi/controllers/admin:GoodsTypeAttrController"],
        beego.ControllerComments{
            Method: "DoEdit",
            Router: `/edit`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["xiaomi/controllers/admin:GoodsTypeController"] = append(beego.GlobalControllerRouter["xiaomi/controllers/admin:GoodsTypeController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["xiaomi/controllers/admin:GoodsTypeController"] = append(beego.GlobalControllerRouter["xiaomi/controllers/admin:GoodsTypeController"],
        beego.ControllerComments{
            Method: "Add",
            Router: `/add`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["xiaomi/controllers/admin:GoodsTypeController"] = append(beego.GlobalControllerRouter["xiaomi/controllers/admin:GoodsTypeController"],
        beego.ControllerComments{
            Method: "DoAdd",
            Router: `/add`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["xiaomi/controllers/admin:GoodsTypeController"] = append(beego.GlobalControllerRouter["xiaomi/controllers/admin:GoodsTypeController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/delete`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["xiaomi/controllers/admin:GoodsTypeController"] = append(beego.GlobalControllerRouter["xiaomi/controllers/admin:GoodsTypeController"],
        beego.ControllerComments{
            Method: "Edit",
            Router: `/edit`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["xiaomi/controllers/admin:GoodsTypeController"] = append(beego.GlobalControllerRouter["xiaomi/controllers/admin:GoodsTypeController"],
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

    beego.GlobalControllerRouter["xiaomi/controllers/admin:NavController"] = append(beego.GlobalControllerRouter["xiaomi/controllers/admin:NavController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["xiaomi/controllers/admin:NavController"] = append(beego.GlobalControllerRouter["xiaomi/controllers/admin:NavController"],
        beego.ControllerComments{
            Method: "Add",
            Router: `/add`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["xiaomi/controllers/admin:NavController"] = append(beego.GlobalControllerRouter["xiaomi/controllers/admin:NavController"],
        beego.ControllerComments{
            Method: "DoAdd",
            Router: `/add`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["xiaomi/controllers/admin:NavController"] = append(beego.GlobalControllerRouter["xiaomi/controllers/admin:NavController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/delete`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["xiaomi/controllers/admin:NavController"] = append(beego.GlobalControllerRouter["xiaomi/controllers/admin:NavController"],
        beego.ControllerComments{
            Method: "Edit",
            Router: `/edit`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["xiaomi/controllers/admin:NavController"] = append(beego.GlobalControllerRouter["xiaomi/controllers/admin:NavController"],
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
            Method: "Auth",
            Router: `/auth`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["xiaomi/controllers/admin:RoleController"] = append(beego.GlobalControllerRouter["xiaomi/controllers/admin:RoleController"],
        beego.ControllerComments{
            Method: "DoAuth",
            Router: `/auth`,
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

    beego.GlobalControllerRouter["xiaomi/controllers/admin:SettingController"] = append(beego.GlobalControllerRouter["xiaomi/controllers/admin:SettingController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["xiaomi/controllers/admin:SettingController"] = append(beego.GlobalControllerRouter["xiaomi/controllers/admin:SettingController"],
        beego.ControllerComments{
            Method: "DoEdit",
            Router: `/edit`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
