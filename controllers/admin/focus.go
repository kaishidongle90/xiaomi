package admin

import (
	"fmt"
	"strconv"
	"xiaomi/models"
)

type FocusController struct {
	BaseController
}

// @router / [get]
func (f *FocusController) Get()  {
	focus := []models.Focus{}
	models.DB.Find(&focus)
	f.Data["focusList"] = focus
	f.TplName = "admin/focus/index.html"
}

// @router /add [get]
func (f *FocusController) Add()  {
	f.TplName = "admin/focus/add.html"
}

// @router /add [post]
func (f *FocusController) DoAdd()  {
	focusType, err1 := f.GetInt("focus_type")
	title := f.GetString("title")
	link := f.GetString("link")
	sort, err2 := f.GetInt("sort")
	status, err3 := f.GetInt("status")

	if err1 != nil || err3 != nil {
		f.Error("非法请求", "/admin/focus/add")
	}
	if err2 != nil {
		f.Error("排序表单里面输入的数据不合法", "/admin/focus/add")
	}
	picPath, err := f.UploadImg("focus_img")
	if err != nil{
		fmt.Println("上传图片错误",err)
		f.Error("上传图片错误","/admin/focus/add")
		return
	}else {
		focus := models.Focus{
			Title:     title,
			FocusType: focusType,
			FocusImg:  picPath,
			Link:      link,
			Sort:      sort,
			Status:    status,
			AddTime:   int(models.GetUnix()),
		}
		models.DB.Create(&focus)
		f.Success("增加轮播图成功", "/admin/focus")
	}
}


// @router /edit [get]
func (f *FocusController) Edit()  {
	id, err := f.GetInt("id")
	if err != nil{
		f.Error("获取轮播图ID失败","/admin/focus/add")
	}
	focus := models.Focus{Id:id}
	models.DB.Find(&focus)
	f.Data["focus"] = focus
	f.TplName = "admin/focus/edit.html"
}

// @router /edit [post]
func (f *FocusController) DOEdit() {
	id, err1 := f.GetInt("id")
	focusType, err2 := f.GetInt("focus_type")
	title := f.GetString("title")
	link := f.GetString("link")
	sort, err3 := f.GetInt("sort")
	status, err4 := f.GetInt("status")

	if err1 != nil || err2 != nil || err4 != nil {
		f.Error("非法请求", "/admin/focus")
	}
	if err3 != nil {
		f.Error("排序表单里面输入的数据不合法", "/admin/focus/edit?id="+strconv.Itoa(id))
	}
	//执行图片上传
	focusImgSrc, err := f.UploadImg("focus_img")
	if err != nil {
		f.Error("上传图片错误", "/admin/focus/edit?id="+strconv.Itoa(id))
	}
	focus := models.Focus{Id: id}
	models.DB.Find(&focus)
	focus.Title = title
	focus.FocusType = focusType
	focus.Link = link
	focus.Sort = sort
	focus.Status = status
	if focusImgSrc != "" {
		focus.FocusImg = focusImgSrc
	}
	err5 := models.DB.Save(&focus).Error
	if err5 != nil {
		f.Error("修改数据失败", "/focus/edit?id="+strconv.Itoa(id))
		return
	}
	f.Success("修改数据成功", "/admin/focus")
}

// @router /delete [get]
func (f *FocusController) Delete() {
	id, err := f.GetInt("id")
	if err != nil{
		f.Error("获取轮播图ID失败","/admin/focus")
	}
	focus := models.Focus{Id:id}
	models.DB.Delete(&focus)
	f.Success("删除轮播图成功", "/admin/focus")
}