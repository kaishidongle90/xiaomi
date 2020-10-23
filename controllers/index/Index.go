package index

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/jinzhu/gorm"
	"strings"
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
	//左侧分类  https://gorm.io/zh_CN/docs/preload.html
	goodsCate := []models.GoodsCate{}
	models.DB.Preload("GoodsCateItem", func(db *gorm.DB) *gorm.DB {
		return db.Where("goods_cate.status=1").Order("goods_cate.sort DESC")
	}).Where("pid=0 AND status=1").Order("sort desc").Find(&goodsCate)
	c.Data["goodsCateList"] = goodsCate
	//获取中间导航
	midNav := []models.Nav{}
	models.DB.Where("status=1 AND Position=2").Order("sort desc").Find(&midNav)
	for i:=0;i<len(midNav);i++{
		midGoods := []models.Goods{}
		relationId := strings.Split(midNav[i].Relation,",")
		models.DB.Where("id in (?)",relationId ).Select("id,title,goods_img,price").Find(&midGoods)
		midNav[i].GoodsItem = midGoods
	}
	c.Data["middleNavList"] = midNav
	//获取楼层数据
	//手机
	phone := models.GetGoodsByCategory(1, "hot", 8)
	c.Data["phoneList"] = phone

	//电视
	tv := models.GetGoodsByCategory(4, "best", 8)
	c.Data["tvList"] = tv
	fmt.Println(tv,phone)
	c.TplName = "index/index/index.html"
}
