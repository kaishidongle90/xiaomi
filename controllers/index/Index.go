package index

import (
	"github.com/astaxie/beego"
	"xiaomi/models"
)

type IndexController struct {
	beego.Controller
}

func (c *IndexController) Get() {
	//获取顶部导航
	topNav := []models.Nav{}
	models.DB.Where("status=1 AND position=1").Find(&topNav)
	c.Data["topNavList"] = topNav
	//获取轮播图
	focus := []models.Focus{}
	models.DB.Where("status=1 AND focus_type=1").Find(&focus)
	c.Data["focusList"] = focus
	//左侧分类

	//...
	c.TplName = "index/index/index.html"
}
