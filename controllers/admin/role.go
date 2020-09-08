package admin

import (
	"strconv"
	"strings"
	"xiaomi/models"
)

type RoleController struct {
	 BaseController
}
// @router  / [get]
func (c *RoleController) Get() {
	roles := []models.Role{}
	models.DB.Find(&roles)
	c.Data["roleList"] = roles
	//fmt.Println("roles:", roles)
	c.TplName = "admin/role/index.html"
}

// @router  /add [get]
func (c *RoleController) Add() {
	c.TplName = "admin/role/add.html"
}
// @router  /add [post]
func (c *RoleController) DoAdd() {
	title := strings.Trim(c.GetString("title")," ")
	description := strings.Trim(c.GetString("description")," ")
	var role models.Role
	role.Title = title
	role.Description = description
	role.Status = 1
	role.AddTime = int(models.GetUnix())
	err := models.DB.Create(&role).Error
	if err != nil {
		c.Error("增加角色", "/admin/role/add")
	} else {
		c.Success("增加角色成功", "/admin/role")
	}
	//c.Ctx.WriteString("执行增加权限")
}

// @router  /edit [get]
func (c *RoleController) Edit(){
	id ,err  := c.GetInt("id")
	if err != nil {
		c.Error("id不存在","/admin/role")
	}
	role := &models.Role{Id:id}
	models.DB.Find(&role)
	c.Data["role"] = role
	c.TplName = "admin/role/edit.html"
}

// @router  /edit [post]
func (c *RoleController) DoEdit() {
	id ,err  := c.GetInt("id")
	if err != nil {
		c.Error("id不存在","/admin/role")
	}
	title := strings.Trim(c.GetString("title")," ")
	description := strings.Trim(c.GetString("description")," ")
	if title == ""{
		c.Error("title不能为空","/admin/role")
		return
	}
	role := models.Role{Id:id}
	role.Title =title
	role.Description = description
	err2 := models.DB.Save(&role).Error
	if err2 != nil{
		c.Error("修改数据失败", "/role/edit?id="+strconv.Itoa(id))
	} else {
		c.Success("修改角色成功", "/admin/role")
	}
}

// @router  /delete [get]
func (c *RoleController) Delete() {
	id ,err  := c.GetInt("id")
	if err != nil {
		c.Error("id不存在","/admin/role")
	}
	role := models.Role{Id:id}
	models.DB.Delete(&role)
	c.Success("删除角色成功", "/admin/role")
}