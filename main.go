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
	//注册模板函数
	beego.AddFuncMap("unixToDate", models.UnixToDate)

	beego.Run()
	defer models.DB.Close()
}

