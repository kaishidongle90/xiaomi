package admin

import (
	"github.com/astaxie/beego"
)

type AdminController struct {
	beego.Controller
}

func (this *AdminController) Get()  {
	this.TplName = "admin/manager/login.html"
}


// @router /kai [get]
func (this *AdminController) Index()  {
	this.TplName = "admin/manager/login.html"
}

// @router / [get]
func (this *AdminController) Manager()  {
	this.TplName = "admin/manager/index.html"
}

// @router /add [get]
func (this *AdminController) Add()  {
	this.TplName = "admin/manager/add.html"
}

// @router /edit [get]
func (this *AdminController) Edit()  {
	this.TplName = "admin/manager/edit.html"
}