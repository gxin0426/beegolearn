package controllers

import (
	"fmt"

	model "github.com/gxin0426/beegolearn/models"
)

type HomeController struct {
	BaseController
}

func (this *HomeController) Get() {

	page, _ := this.GetInt("page")
	if page <= 0 {
		page = 1
	}

	var artList []model.Article
	artList = model.FindArticleWithPage(page)

	this.Data["pageCode"] = 1
	this.Data["HasFooter"] = true

	fmt.Println("islogin:", this.IsLogin, this.Loginuser)
	this.Data["Content"] = model.MakeHomeBlocks(artList, this.IsLogin)
	this.TplName = "home.html"
}
