package controllers

import (
	"beegotest/model"
	"strconv"
)

type ShowArticleController struct {
	BaseController
}

func (this *ShowArticleController) Get() {
	idStr := this.Ctx.Input.Param(":id")

	id, _ := strconv.Atoi(idStr)

	art := model.QueryArticleWithId(id)

	this.Data["Title"] = art.Title
	this.Data["Content"] = art.Content
	this.TplName = "show_article.html"
}
