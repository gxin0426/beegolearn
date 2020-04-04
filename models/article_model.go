package model

import (
	"fmt"

	"github.com/astaxie/beego"
	utils "github.com/gxin0426/beegolearn/utils"
)

type Article struct {
	Id         int
	Title      string
	Tags       string
	Short      string
	Content    string
	Author     string
	Createtime int64
}

// --------------数据处理-------------

func AddArticle(article Article) (int64, error) {
	i, err := insertArticle(article)
	return i, err
}

//------------数据库操作--------------

func insertArticle(article Article) (int64, error) {
	return utils.ModifyDB("insert into article(title, tags, short, content, author, createtime) values (?, ?, ?, ?, ?, ?)",
		article.Title, article.Tags, article.Short, article.Content, article.Author, article.Createtime)
}

//查询文章

func FindArticleWithPage(page int) ([]Article, error) {
	num, _ := beego.AppConfig.Int("articleListPageNum")
	page--
	fmt.Println("page:", page)

	return QueryArticleWithPage(page, num), nil
}

func QueryArticleWithPage(page, num int) ([]Article, error) {
	sql := fmt.Sprintf("limit %d, %d", page*num, num)
	return QueryArticlesWithCon(sql), nil
}

func QueryArticlesWithCon(sql string) []Article {
	sql = "select id, title, tags, short, content, author, createtime, from article" + sql
	rows := utils.QueryRowDB(sql)
	var artList []Article
	for rows.Next() {
		id := 0
		title := ""
		tags := ""
		short := ""
		content := ""
		author := ""
		var createtime int64
		createtime = 0
		rows.Scan(&id, &title, &tags, &short, &content, &author, &createtime)
		art := Article{id, title, tags, short, content, author, createtime}
		artList = append(artList, art)
	}
	return artList
}
