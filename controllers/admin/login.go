package admin

import (
	"fmt"
	"xiaomi/models"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	"github.com/astaxie/beego/utils/captcha"
)

var cpt *captcha.Captcha

func init() {
	store := cache.NewMemoryCache()
	cpt = captcha.NewWithFilter("/captcha/", store)
	cpt.ChallengeNums = 4
	cpt.StdWidth = 100
	cpt.StdHeight = 40
}

type LoginController struct {
	BaseController
}

func (c *LoginController) Get() {
	//获取user表的数据验证数据库是否连接成功

	c.TplName = "admin/login/login.html"
}

func (c *LoginController) DoLogin() {

	//1、验证用户输入的验证码是否正确
	var flag = cpt.VerifyReq(c.Ctx.Request)
	if flag {
		//2、获取表单传过来的用户名和密码
		username := strings.Trim(c.GetString("username"), "")
		password := models.Md5(strings.Trim(c.GetString("password"), ""))
		//3、去数据库匹配
		manager := []models.Manager{}
		models.DB.Where("username=? AND password=?", username, password).Find(&manager)
		fmt.Println(manager)
		if len(manager) > 0 {
			//登录成功 1、保存用户信息  2、跳转到后台管理系统
			c.SetSession("userinfo", manager[0])
			c.Success("登录成功", "/"+beego.AppConfig.String("adminPath"))
		} else {
			c.Error("用户名获取密码错误", "/"+beego.AppConfig.String("adminPath")+"/login")
		}

	} else {
		c.Error("验证码错误", "/"+beego.AppConfig.String("adminPath")+"/login")
	}
}


func (c *LoginController) LoginOut() {
	c.DelSession("userinfo")
	c.Success("退出登录成功", "/admin/login")
}