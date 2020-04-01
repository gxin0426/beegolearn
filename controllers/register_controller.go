package controllers

import (
	"fmt"
	"time"

	"github.com/astaxie/beego"
	"github.com/gxin0426/beegolearn/model"
	"github.com/gxin0426/beegolearn/utils"
)


type RegisterController struct {
	beego.Controller
}

func (this *RegisterController) Get() {
	this.TplName = "register.html"
}

func (this *RegisterController) Post() {

	//获取表单信息
	username := this.GetString("username")
	password := this.GetString("password")
	repassword := this.GetString("repassword")

	fmt.Println(username, password, repassword)

	//log.INFO(username, password, repassword)

	// 注册前先判断该用户是否存在 如果已经存在返回错误

	id := model.QueryUserWithUsername(username)

	fmt.Println("id:", id)

	if id > 0 {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "用户已存在"}
		this.ServeJSON()
		return
	}

	//注册用户名密码
	//存储密码是md5后的数据 那么在登录验证的时候 也是需要将用户的密码md5 之后和数据库密码进行判断

	password = utils.MD5(password)

	user := model.User{0, username, password, 0, time.Now().Unix()}

	_, err := model.InsertUser(user)

	if err != nil {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "failure"}

	} else {
		this.Data["json"] = map[string]interface{}{"code": 1, "message": "success"}
	}
	this.ServeJSON()

}
