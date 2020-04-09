package controllers

import "beegotest/model"

type TagsController struct {
	BaseController
}

func (this *TagsController) Get() {
	tags := model.QueryArticleWithParam("tags")

	this.Data["Tags"] = model.HandleTagsListData(tags)

	this.TplName = "tags.html"

}
