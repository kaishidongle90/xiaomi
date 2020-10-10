package main

import (
	"encoding/gob"
	"github.com/astaxie/beego"
	_ "github.com/astaxie/beego/session/redis"
	"xiaomi/models"
	_ "xiaomi/routers"
)

func init() {
	gob.Register(models.Manager{})
}
func main() {
	beego.BConfig.WebConfig.Session.SessionProvider = "redis"
	beego.BConfig.WebConfig.Session.SessionProviderConfig = "127.0.0.1:6379"
	//注册模板函数
	beego.AddFuncMap("unixToDate", models.UnixToDate)
	beego.Run()
	defer models.DB.Close()
}

