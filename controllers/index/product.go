package index

import (
	"math"
	"strconv"
	"xiaomi/models"
)

type ProductController struct {
	BaseController
}

func (c *ProductController) CategoryList() {
	//调用公共功能
	c.SuperInit()

	id := c.Ctx.Input.Param(":id")
	cateId, _ := strconv.Atoi(id)
	curretGoodsCate := models.GoodsCate{}
	subGoodsCate := []models.GoodsCate{}
	models.DB.Where("id=?", cateId).Find(&curretGoodsCate)
	//分页
	//当前页
	page, _ := c.GetInt("page")
	if page == 0 {
		page = 1
	}
	//每一页显示的数量
	pageSize := 5

	var tempSlice []int
	if curretGoodsCate.Pid == 0 { //顶级分类
		//二级分类
		models.DB.Where("pid=?", curretGoodsCate.Id).Find(&subGoodsCate)
		for i := 0; i < len(subGoodsCate); i++ {
			tempSlice = append(tempSlice, subGoodsCate[i].Id)
		}
	} else {
		//获取当前二级分类对应的兄弟分类
		models.DB.Where("pid=?", curretGoodsCate.Pid).Find(&subGoodsCate)
	}
	tempSlice = append(tempSlice, cateId)
	where := "cate_id in (?)"
	goods := []models.Goods{}
	models.DB.Where(where, tempSlice,).Select("id,title,price,goods_img,sub_title").Offset((page - 1) * pageSize).Limit(pageSize).Order("sort desc").Find(&goods)
	//查询goods表里面的数量
	var count int
	models.DB.Where(where, tempSlice).Table("goods").Count(&count)

	c.Data["goodsList"] = goods
	c.Data["subGoodsCate"] = subGoodsCate
	c.Data["curretGoodsCate"] = curretGoodsCate
	c.Data["totalPages"] = math.Ceil(float64(count) / float64(pageSize))
	c.Data["page"] = page

	//指定分类模板
	tpl := curretGoodsCate.Template
	if tpl == "" {
		tpl = "index/product/list.html"
	}

	c.TplName = tpl
}

