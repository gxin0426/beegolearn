package controllers

import (
	"beegotest/model"
	"log"
)

type DeleteArticleController struct {
	BaseController
}

func (this *DeleteArticleController) Get() {
	id, _ := this.GetInt("id")

	_, err := model.DeleteArticle(id)

	if err != nil {
		log.Println(err)
	}
	this.Redirect("/", 302)
}
