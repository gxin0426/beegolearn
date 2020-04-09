package controllers

import (
	"beegotest/model"
)

type UpdateArticleController struct {
	BaseController
}

func (this *UpdateArticleController) Get() {
	id, _ := this.GetInt("id")
	//fmt.Println(id)
	art := model.QueryArticleWithId(id)
	//fmt.Println(art)
	this.Data["Title"] = art.Title
	this.Data["Tags"] = art.Tags
	this.Data["Short"] = art.Short
	this.Data["Content"] = art.Content
	this.Data["Id"] = art.Id
	this.TplName = "write_article.html"

}

func (this *UpdateArticleController) Post() {
	id, _ := this.GetInt("id")

	title := this.GetString("title")
	tags := this.GetString("tags")
	short := this.GetString("short")
	content := this.GetString("content")

	article := model.Article{id, title, tags, short, content, "", 0}

	_, err := model.UpdateArticle(article)

	//返回数据给浏览器
	if err == nil {
		this.Data["json"] = map[string]interface{}{"code": 1, "message": "success"}
	} else {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "failure"}
	}
	this.ServeJSON()
}
