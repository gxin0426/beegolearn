package controllers

import (
	"fmt"
	"time"

	model "github.com/gxin0426/beegolearn/models"
)

type AddArticleController struct {
	BaseController
}

//当访问/add路径的时候会触发AddarticleController 的Get方法 响应的页面是通过TplName

func (this *AddArticleController) Get() {
	this.TplName = "write_article.html"
}

//通过this.ServerJSON()这个方法返回json字符串
func (this *AddArticleController) Post() {

	//获取浏览器传输的数据 通过表单的name属性获取值
	title := this.GetString("title")
	tags := this.GetString("tags")
	short := this.GetString("short")
	content := this.GetString("content")
	fmt.Printf("title: %s", title)

	res := model.Article{0, title, tags, short, content, "gree", time.Now().Unix()}

	_, err := model.AddArticle(res)

	var resp map[string]interface{}

	if err == nil {
		resp = map[string]interface{}{"code": 1, "message": "ok"}
	} else {
		resp = map[string]interface{}{"code": 0, "message": "error"}
	}

	this.Data["json"] = resp

	this.ServeJSON()

}
