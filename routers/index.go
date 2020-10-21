package routers
import (
	"xiaomi/controllers/index"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &index.IndexController{})
	beego.Router("/product", &index.ProductController{})
	//beego.Router("/user", &index.UserController{})

}
