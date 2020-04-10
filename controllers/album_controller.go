package controllers

import (
	"beegotest/model"
	"fmt"
)

type AlbumController struct {
	BaseController
}

func (this *AlbumController) Get() {
	albums, err := model.FindAllAlbums()
	if err != nil {
		fmt.Println(err)
	}
	this.Data["Album"] = albums
	this.TplName = "album.html"
}
