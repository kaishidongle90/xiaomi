package admin

import (
	"strconv"
	"strings"
	"xiaomi/models"
)

type ManagerController struct {
	BaseController
}

// @router / [get]
func (this *ManagerController) Manager() {
	managerList := []models.Manager{}
	models.DB.Preload("Role").Find(&managerList)
	this.Data["managerList"] = managerList
	this.TplName = "admin/manager/index.html"
}

// @router /add [get]
func (this *ManagerController) Add() {
	roleList := []models.Role{}
	models.DB.Find(&roleList)
	this.Data["roleList"] = roleList
	this.TplName = "admin/manager/add.html"
}

// @router /add [post]
func (this *ManagerController) DoAdd()  {
	role_id, err:= this.GetInt("role_id")
	if err != nil{
		this.Error("获取角色ID失败","/admin/manager/add")
	}
	username := strings.Trim(this.GetString("username"), " ")
	password := strings.Trim(this.GetString("password"), " ")
	mobile := strings.Trim(this.GetString("mobile"), " ")
	email := strings.Trim(this.GetString("email"), " ")

	if len(username) < 2 || len(password) < 6 {
		this.Error("用户名或者密码长度不合法", "/admin/manager/add")
		return
	}
	//判断数据库里面有没有当前用户
	users := []models.Manager{}
	models.DB.Where("name = ?",username).Find(&users)
	if len(users) >0 {
		this.Error("用户已存在", "/admin/manager/add")
		return
	}

	user := models.Manager{}
	user.Username = username
	user.Password = models.Md5(password)
	user.Mobile, _= strconv.Atoi(mobile)
	user.Email = email
	user.Status = 1
	user.AddTime = int(models.GetUnix())
	user.RoleId = role_id
	err1 := models.DB.Create(&user).Error
	if err1 != nil {
		this.Error("增加管理员失败", "/admin/manager/add")
		return
	}
	this.Success("增加管理员成功", "/admin/manager")

}


// @router /edit [get]
func (this *ManagerController) Edit() {
	this.TplName = "admin/manager/edit.html"
}
