package admin

import (
	"fmt"
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

// @router /auth [get]
func (this *RoleController) Auth()  {
	//获取角色ID
	roleId ,err  := this.GetInt("id")
	if err != nil {
		this.Error("roleId不存在","/admin/role")
	}
	this.Data["roleId"] = roleId

	// 获取所有权限
	access := []models.Access{}
	models.DB.Preload("AccessItem" ).Where("module_id = 0").Find(&access)

	//fmt.Println(access)

  // 获取当前用户的所有权限
   roleAccess := []models.RoleAccess{}
   models.DB.Where("role_id = ?", roleId).Find(&roleAccess)
   roleAccessMap := make(map[int]int)
   for _, v  := range roleAccess{
   	  roleAccessMap[v.AccessId] = v.AccessId
   }
	fmt.Println(roleAccessMap)
   //循环遍历所有的权限数据，判断当前权限的id是否在角色权限的Map对象中,如果是的话给当前数据加入checked属性
	for k,v := range access{
		if _, ok := roleAccessMap[v.Id]; ok{
			fmt.Println("=====")
			access[k].Checked = true
		}
		for j, ai := range v.AccessItem{
			if _, ok := roleAccessMap[ai.Id]; ok{
				access[k].AccessItem[j].Checked = true
			}
		}

	}
	fmt.Println(access)
	this.Data["accessList"] = access
	this.TplName = "admin/role/auth.html"
}

// @router /auth [post]
func (this *RoleController) DoAuth()  {
	//获取角色ID
	roleId ,err  := this.GetInt("role_id")
	fmt.Println(roleId)
	if err != nil {
		this.Error("roleId不存在","/admin/role")
	}
	//// 获取权限列表
	pris := this.GetStrings("access_node")

	//清空已有权限列表
	ra := models.RoleAccess{}
	models.DB.Where("role_id = ?", roleId).Delete(&ra)

	//新增权限
	for i:=0;i<len(pris);i++{
		ra := models.RoleAccess{}
		ra.RoleId = roleId
		ra.AccessId,_ = strconv.Atoi(pris[i])
		models.DB.Create(&ra)
	}
	this.Success("授权成功", "/admin/role/auth?id="+strconv.Itoa(roleId))
}