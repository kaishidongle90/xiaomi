package routers
import (
	"xiaomi/controllers/index"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &index.IndexController{})
	beego.Router("/category_:id([0-9]+).html", &index.ProductController{}, "get:CategoryList")

}
