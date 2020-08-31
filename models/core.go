package models

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB
var err error

func init() {
	mysqladmin := beego.AppConfig.String("mysqladmin")
	mysqlpwd := beego.AppConfig.String("mysqlpwd")
	mysqldb := beego.AppConfig.String("mysqldb")

	DB, err = gorm.Open("mysql", mysqladmin+":"+mysqlpwd+"@/"+mysqldb+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("链接mysql数据库失败")
	}
}
