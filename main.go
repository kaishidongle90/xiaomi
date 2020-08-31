package main

import (
	"github.com/astaxie/beego"
	"xiaomi/models"
	_ "xiaomi/routers"
)

func main() {
	beego.Run()
	defer models.DB.Close()
}
