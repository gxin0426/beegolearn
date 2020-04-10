package model

import (
	"bytes"
	"fmt"
	"html/template"
	"strconv"
	"strings"

	"beegotest/utils"

	"github.com/astaxie/beego"
)

type HomeBlockParam struct {
	Id         int
	Title      string
	Tags       []TagLink
	Short      string
	Content    string
	Author     string
	CreateTime string
	//查看文章地址
	Link string
	//修改文章的地址
	UpdateLink string
	DeleteLink string

	//记录是否登录
	IsLogin bool
}

//分页结构体

type HomeFooterPageCode struct {
	HasPre   bool
	HasNext  bool
	ShowPage string
	PreLink  string
	NextLink string
}

type TagLink struct {
	TagName string
	TagUrl  string
}

func MakeHomeBlocks(articles []Article, isLogin bool) template.HTML {
	htmlHome := ""

	for _, art := range articles {
		homeParam := HomeBlockParam{}
		homeParam.Id = art.Id
		homeParam.Title = art.Title
		homeParam.Tags = createTagsLink(art.Tags)
		homeParam.Short = art.Short
		homeParam.Content = art.Content
		homeParam.Author = art.Author
		homeParam.CreateTime = utils.SwitchTimeStampToData(art.Createtime)
		homeParam.Link = "/article/" + strconv.Itoa(art.Id)
		homeParam.DeleteLink = "/article/delete?id=" + strconv.Itoa(art.Id)
		homeParam.UpdateLink = "/article/update?id=" + strconv.Itoa(art.Id)
		homeParam.IsLogin = isLogin

		//处理变量
		//ParseFile解析该文件 用于插入变量
		t, _ := template.ParseFiles("views/block/home_block.html")
		buffer := bytes.Buffer{}
		//就是将html文件里面的比那两替换为传进去的数据
		t.Execute(&buffer, homeParam)
		htmlHome += buffer.String()
	}
	//fmt.Println("htmlHome-->", template.HTML(htmlHome))
	return template.HTML(htmlHome)
}

func createTagsLink(tags string) []TagLink {

	var tagLink []TagLink
	tagsParam := strings.Split(tags, ",")
	for _, tag := range tagsParam {
		tagLink = append(tagLink, TagLink{tag, "/tag=" + tag})
	}
	return tagLink
}

//翻页功能

func ConfigHomeFooterPageCode(page int) HomeFooterPageCode {
	pageCode := HomeFooterPageCode{}

	//查询出总的条数
	num := GetArticleRowsNum()

	//从配置文件中读取每页显示的条数
	pageRow, _ := beego.AppConfig.Int("articleListPageNum")

	//计算出总页数
	fmt.Println("zongtiaoshu : ", num)
	fmt.Println("kkk", pageRow)
	allPageNum := (num-1)/pageRow + 1
	fmt.Println("yeshu", allPageNum)
	fmt.Printf("%d%d\n", page, allPageNum)
	pageCode.ShowPage = fmt.Sprintf("%d/%d", page, allPageNum)

	//当前页数小于等于1， 那么上一页的按钮不能点击
	if page <= 1 {
		pageCode.HasPre = false
	} else {
		pageCode.HasPre = true
	}

	//当前页数大于等于总页数， 那么下一页的按钮不能点击

	if page >= allPageNum {
		pageCode.HasNext = false
	} else {
		pageCode.HasNext = true
	}

	pageCode.PreLink = "/?page=" + strconv.Itoa(page-1)
	pageCode.NextLink = "/?page=" + strconv.Itoa(page+1)
	return pageCode

}
