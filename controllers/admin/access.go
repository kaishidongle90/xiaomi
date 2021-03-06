package admin

import (
	"fmt"
	"strconv"
	"xiaomi/models"
)

type AccessController struct {
	BaseController
}
// @router / [get]
func (this *AccessController) Get() {
	access :=[]models.Access{}
	//models.DB.Find(&access)
	models.DB.Preload("AccessItem").Where("module_id=0").Find(&access)
	//fmt.Println(access)
	this.Data["accessList"] = access
	this.TplName = "admin/access/index.html"
}

// @router /add [get]
func (this *AccessController) Add(){
	access :=[]models.Access{}
	models.DB.Where("module_id = 0").Find(&access)
	this.Data["accessList"] = access
	this.TplName = "admin/access/add.html"
}

// @router /add [post]
func (this *AccessController) DoAdd() {
	moduleName := this.GetString("module_name")
	iType, err1 := this.GetInt("type")
	actionName := this.GetString("action_name")
	url := this.GetString("url")
	moduleId, err2 := this.GetInt("module_id")
	sort, err3 := this.GetInt("sort")
	description := this.GetString("description")
	status, err4 := this.GetInt("status")
	if err1 != nil || err2 != nil || err3 != nil || err4 != nil {
		this.Error("传入参数错误", "/access/add")
		return
	}
	access := models.Access{
		ModuleName:  moduleName,
		Type:        iType,
		ActionName:  actionName,
		Url:         url,
		ModuleId:    moduleId,
		Sort:        sort,
		Description: description,
		Status:      status,
	}
	err := models.DB.Create(&access).Error
	if err != nil {
		this.Error("增加数据失败", "/admin/access/add")
	} else {
		this.Success("增加数据成功", "/admin/access")
	}
}
// @router /edit [get]
func (c *AccessController) Edit() {
	//获取要修改的数据
	id, err1 := c.GetInt("id")
	if err1 != nil {
		c.Error("传入参数错误", "/access")
		return
	}
	access := models.Access{Id: id}
	models.DB.Find(&access)
	c.Data["access"] = access

	//获取顶级模块
	accessList := []models.Access{}
	models.DB.Where("module_id=0").Find(&accessList)
	c.Data["accessList"] = accessList

	c.TplName = "admin/access/edit.html"
}

// @router /edit [post]
func (c *AccessController) DoEdit() {
	id, err1 := c.GetInt("id")
	moduleName := c.GetString("module_name")
	iType, err2 := c.GetInt("type")
	actionName := c.GetString("action_name")
	url := c.GetString("url")
	moduleId, err3 := c.GetInt("module_id")
	sort, err4 := c.GetInt("sort")
	description := c.GetString("description")
	status, err5 := c.GetInt("status")
	if err1 != nil || err2 != nil || err3 != nil || err4 != nil || err5 != nil {
		c.Error("传入参数错误", "/access")
		return
	}
	access := models.Access{Id: id}
	models.DB.Find(&access)
	access.ModuleName = moduleName
	access.Type = iType
	access.ActionName = actionName
	access.Url = url
	access.ModuleId = moduleId
	access.Sort = sort
	access.Description = description
	access.Status = status
	err := models.DB.Save(&access).Error
	if err != nil {
		c.Error("修改失败", "/access/edit?id="+strconv.Itoa(id))
		return
	}
	c.Success("修改成功", "/access/")
}

// @router /delete [get]
func (this *AccessController) Delete() {
	id, err := this.GetInt("id")
	if err != nil{
		this.Error("传入参数错误", "/admin/access")
		return
	}
	access := models.Access{Id:id}
	models.DB.Find(&access)
	if access.ModuleId == 0 {
		as := []models.Access{}
		models.DB.Debug().Where("module_id = ?",access.Id).Find(&as)
		if len(as) > 0 {
			fmt.Println("++++++")
			this.Error("模块下面还有数据不能删除", "/admin/access")
			return
		}
	}
	models.DB.Debug().Delete(&access)
	this.Success("删除成功", "/admin/access")
}