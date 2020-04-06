package model

import (
	"fmt"
	"strconv"

	utils "beegotest/utils"

	"github.com/astaxie/beego"
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
	SetArticleRowsNum()
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
	fmt.Println("aaaaaaaaaaaaaaaaaa")
	return QueryArticleWithPage(page, num)
}

func QueryArticleWithPage(page, num int) ([]Article, error) {
	sql := fmt.Sprintf(" limit %d,%d", page*num, num)
	fmt.Println("BBBBBBBBBBB")
	return QueryArticlesWithCon(sql)
}

func QueryArticlesWithCon(sql string) ([]Article, error) {
	sql = "select id, title, tags, short, content, author, createtime from article" + sql
	fmt.Println(sql)
	rows, err := utils.QueryDB(sql)

	if err != nil {
		fmt.Println(err)
		fmt.Println("有错误")
		return nil, err
	}

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
	fmt.Println("articlelist res:", artList)
	return artList, nil
}

//----翻页功能

//存储表的行数 只有自己可以更改 当文章新增或者删除时需要更新这个值

var articleRowsNum = 0

// 只有首次获取行数的时候才去统计表里的行数
func GetArticleRowsNum() int {
	if articleRowsNum == 0 {
		articleRowsNum = QueryArticleRowNum()
	}
	return articleRowsNum
}

func QueryArticleRowNum() int {
	row := utils.QueryRowDB("select count(id) from article")

	num := 0

	row.Scan(&num)
	fmt.Println("yigong :", num)
	return num
}

//设置页数
func SetArticleRowsNum() {
	articleRowsNum = QueryArticleRowNum()
}

//----------按照标签查询-------------

func QueryArticleWithTag(tag string) ([]Article, error) {
	sql := "where tags like '%," + tag + ",%'"
	sql += " or tags like '%," + tag + "'"
	sql += " or tags like '" + tag + ",%'"
	sql += " or tags like '" + tag + "'"

	return QueryArticlesWithCon(sql)
}

func QueryArticleWithId(id int) Article {
	row := utils.QueryRowDB("select id, title, tags, content, author, createtime from article where id =" + strconv.Itoa(id))
	title := ""
	tags := ""
	short := ""
	content := ""
	author := ""
	var createtime int64
	createtime = 0

	row.Scan(&id, &title, &tags, &short, &content, &author, &createtime)
	art := Article{id, title, tags, short, content, author, createtime}
	return art
}
