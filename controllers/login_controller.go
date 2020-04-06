package controllers

import (
	"beegotest/model"
	utils "beegotest/utils"

	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}

func (this *LoginController) Get() {
	this.TplName = "login.html"

}

func (this *LoginController) Post() {
	username := this.GetString("username")
	password := this.GetString("password")

	id := model.QueryUserWithParam(username, utils.MD5(password))

	if id > 0 {
		this.SetSession("loginuser", username)
		this.Data["json"] = map[string]interface{}{"code": 1, "message": "login success"}
	} else {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "login failure"}
	}

	this.ServeJSON()

}
