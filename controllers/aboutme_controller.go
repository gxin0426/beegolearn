package controllers

import "github.com/astaxie/beego"

type AboutmeController struct {
	beego.Controller
}

func (this *AboutmeController) Get() {
	this.Data["wechat"] = "8888-8888"
	this.Data["qq"] = "8888-6666"
	this.Data["tel"] = "8888-9999"

	this.TplName = "aboutme.html"
}
