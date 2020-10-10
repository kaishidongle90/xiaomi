package admin

import (
	"xiaomi/models"
	"strconv"
)

type GoodsCateController struct {
	BaseController
}
// @router / [get]
func (c *GoodsCateController) Get() {
	goodsCate := []models.GoodsCate{}
	models.DB.Preload("GoodsCateItem").Where("pid=0").Find(&goodsCate)
	c.Data["goodsCateList"] = goodsCate
	c.TplName = "admin/goodsCate/index.html"
}

// @router /add
func (c *GoodsCateController) Add() {
	//加载顶级模块
	goodsCate := []models.GoodsCate{}
	models.DB.Where("pid=0").Find(&goodsCate)
	c.Data["goodsCateList"] = goodsCate
	c.TplName = "admin/goodsCate/add.html"
}
// @router /add [post]
func (c *GoodsCateController) DoAdd() {
	title := c.GetString("title")
	pid, err1 := c.GetInt("pid")
	link := c.GetString("link")
	template := c.GetString("template")
	subTitle := c.GetString("sub_title")
	keywords := c.GetString("keywords")
	description := c.GetString("description")
	sort, err2 := c.GetInt("sort")
	status, err3 := c.GetInt("status")

	if err1 != nil || err3 != nil {
		c.Error("传入参数类型不正确", "/goodsCate/add")
		return
	}
	if err2 != nil {
		c.Error("排序值必须是整数", "/goodsCate/add")
		return
	}
	uploadDir, _ := c.UploadImg("cate_img")
	goodsCate := models.GoodsCate{
		Title:       title,
		Pid:         pid,
		SubTitle:    subTitle,
		Link:        link,
		Template:    template,
		Keywords:    keywords,
		Description: description,
		CateImg:     uploadDir,
		Sort:        sort,
		Status:      status,
		AddTime:     int(models.GetUnix()),
	}
	err := models.DB.Create(&goodsCate).Error
	if err != nil {
		c.Error("增加失败", "/admin/goodscate/add")
		return
	}
	c.Success("增加成功", "/admin/goodscate")

}

// @router /edit [get]
func (c *GoodsCateController) Edit() {

	//获取当前数据
	id, err := c.GetInt("id")
	if err != nil {
		c.Error("传入参数错误", "/admin/goodscate")
		return
	}
	goodsCate := models.GoodsCate{Id: id}
	models.DB.Find(&goodsCate)
	c.Data["goodsCate"] = goodsCate

	//顶级分类
	topGoodsCate := []models.GoodsCate{}
	models.DB.Where("pid=0").Find(&topGoodsCate)
	c.Data["goodsCateList"] = topGoodsCate

	//获取要修改的数据
	c.TplName = "admin/goodsCate/edit.html"
}
// @router /edit [post]
func (c *GoodsCateController) DoEdit() {
	id, err1 := c.GetInt("id")
	title := c.GetString("title")
	pid, err2 := c.GetInt("pid")
	link := c.GetString("link")
	template := c.GetString("template")
	subTitle := c.GetString("sub_title")
	keywords := c.GetString("keywords")
	description := c.GetString("description")
	sort, err3 := c.GetInt("sort")
	status, err4 := c.GetInt("status")

	if err1 != nil || err2 != nil || err4 != nil {
		c.Error("传入参数类型不正确", "/admin/goodscate")
		return
	}
	if err3 != nil {
		c.Error("排序值必须是整数", "/admin/goodscate/edit?id="+strconv.Itoa(id))
		return
	}
	uploadDir, _ := c.UploadImg("cate_img")

	goodsCate := models.GoodsCate{Id: id}
	models.DB.Find(&goodsCate)
	goodsCate.Title = title
	goodsCate.Pid = pid
	goodsCate.Link = link
	goodsCate.Template = template
	goodsCate.SubTitle = subTitle
	goodsCate.Keywords = keywords
	goodsCate.Description = description
	goodsCate.Sort = sort
	goodsCate.Status = status
	if uploadDir != "" {
		goodsCate.CateImg = uploadDir
	}
	err := models.DB.Save(&goodsCate).Error
	if err != nil {
		c.Error("修改失败", "/admin/goodscate/edit?id="+strconv.Itoa(id))
		return
	}
	c.Success("修改成功", "/admin/goodscate")

}
// @router /delete [get]
func (c *GoodsCateController) Delete() {

	id, err1 := c.GetInt("id")
	if err1 != nil {
		c.Error("传入参数错误", "/amdin/goodscate")
		return
	}
	//获取当前数据
	goodsCate1 := models.GoodsCate{Id: id}
	models.DB.Find(&goodsCate1)
	if goodsCate1.Pid == 0 { //顶级分类
		goodsCate3 := []models.GoodsCate{}
			models.DB.Where("pid=?", goodsCate1.Id).Find(&goodsCate3)
		if len(goodsCate3) > 0 {
			c.Error("当前分类下面还子分类，无法删除", "/admin/goodscate")
			return
		}
	}
	//goodsCate2 := models.GoodsCate{Id: id}
	models.DB.Delete(&goodsCate1)
	c.Success("删除成功", "/admin/goodscate")

}
