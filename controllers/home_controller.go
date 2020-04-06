package controllers

import (
	"beegotest/model"
	"fmt"
)

type HomeController struct {
	BaseController
}

func (this *HomeController) Get() {

	//分页  http://localhost:8080?page=2
	//标签  http://localhost:8080?tag=web

	tag := this.GetString("tag")

	fmt.Println("tag:", tag)

	page, _ := this.GetInt("page")

	var artList []model.Article

	if len(tag) > 0 {
		artList, _ = model.QueryArticleWithTag(tag)
		this.Data["HasFooter"] = false
	} else {
		if page <= 0 {
			page = 1
		}

		artList, _ = model.FindArticleWithPage(page)

		this.Data["PageCode"] = model.ConfigHomeFooterPageCode(page)
		this.Data["HasFooter"] = true
	}
	fmt.Println("islogin:", this.IsLogin, this.Loginuser)
	this.Data["Content"] = model.MakeHomeBlocks(artList, this.IsLogin)
	this.TplName = "home.html"
}
