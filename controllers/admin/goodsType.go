package admin

import (
	"xiaomi/models"
	"strconv"
	"strings"
)

type GoodsTypeController struct {
	BaseController
}
// @router / [get]
func (c *GoodsTypeController) Get() {

	goodsType := []models.GoodsType{}
	models.DB.Find(&goodsType)
	c.Data["goodsTypeList"] = goodsType
	c.TplName = "admin/goodsType/index.html"
}

// @router /add [get]
func (c *GoodsTypeController) Add() {
	c.TplName = "admin/goodsType/add.html"
}
// @router /add [post]
func (c *GoodsTypeController) DoAdd() {
	title := strings.Trim(c.GetString("title"), " ")
	description := strings.Trim(c.GetString("description"), " ")
	status, err1 := c.GetInt("status")
	if err1 != nil {
		c.Error("传入参数不正确", "/admin/goodstype/add")
		return
	}
	if title == "" {
		c.Error("标题不能为空", "/admin/goodstype/add")
		return
	}
	goodsType := models.GoodsType{}
	goodsType.Title = title
	goodsType.Description = description
	goodsType.Status = status
	goodsType.AddTime = int(models.GetUnix())
	err := models.DB.Create(&goodsType).Error
	if err != nil {
		c.Error("增加商品类型失败", "/admin/goodstype/add")
	} else {
		c.Success("增加商品类型成功", "/admin/goodstype")
	}
}
// @router /edit [get]
func (c *GoodsTypeController) Edit() {
	id, err := c.GetInt("id")
	if err != nil {
		c.Error("传入参数错误", "/admin/goodstype")
		return
	}

	goodsType := models.GoodsType{Id: id}
	models.DB.Find(&goodsType)
	c.Data["goodsType"] = goodsType
	c.TplName = "admin/goodsType/edit.html"
}
// @router /edit [post]
func (c *GoodsTypeController) DoEdit() {

	id, err1 := c.GetInt("id")

	title := strings.Trim(c.GetString("title"), " ")
	description := strings.Trim(c.GetString("description"), " ")
	status, err2 := c.GetInt("status")
	if err1 != nil || err2 != nil {
		c.Error("传入参数错误", "/admin/goodstype")
		return
	}

	if title == "" {
		c.Error("标题不能为空", "/role/add")
		return
	}
	//修改
	goodsType := models.GoodsType{Id: id}
	models.DB.Find(&goodsType)
	goodsType.Title = title
	goodsType.Description = description
	goodsType.Status = status

	err3 := models.DB.Save(&goodsType).Error
	if err3 != nil {
		c.Error("修改数据失败", "/admin/goodstype/edit?id="+strconv.Itoa(id))
	} else {
		c.Success("修改数据成功", "/admin/goodstype")
	}

}
// @router /delete [get]
func (c *GoodsTypeController) Delete() {
	id, err1 := c.GetInt("id")
	if err1 != nil {
		c.Error("传入参数错误", "/admin/goodstype")
		return
	}
	goodsType := models.GoodsType{Id: id}
	models.DB.Delete(&goodsType)
	c.Success("删除数据成功", "/admin/goodstype")

}
