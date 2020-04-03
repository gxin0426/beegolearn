package utils

import (
	"bytes"
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"fmt"
	"html/template"
	"log"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"github.com/russross/blackfriday"
	"github.com/sourcegraph/syntaxhighlight"
)

var db *sql.DB

func InitMysql() {

	fmt.Println(".... init mysql")

	driverName := beego.AppConfig.String("driveName")

	user := beego.AppConfig.String("mysqluser")
	pwd := beego.AppConfig.String("mysqlpwd")
	host := beego.AppConfig.String("host")
	port := beego.AppConfig.String("port")
	dbname := beego.AppConfig.String("dbname")

	dbConn := user + ":" + pwd + "@tcp(" + host + ":" + port + ")/" + dbname + "?charset=utf8"

	db1, err := sql.Open(driverName, dbConn)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		db = db1

		CreateTableWithUser()

		CreateTableWithArtcle()

	}

}

//操作数据库
func ModifyDB(sql string, args ...interface{}) (int64, error) {
	result, err := db.Exec(sql, args...)
	if err != nil {
		log.Println(err)
		return 0, err
	}

	count, err := result.RowsAffected()
	if err != nil {
		log.Println(err)
		return 0, err
	}
	return count, nil
}

func CreateTableWithUser() {
	sql := `CREATE TABLE IF NOT EXISTS users(
			id INT(4) PRIMARY KEY AUTO_INCREMENT NOT NULL,
			username VARCHAR(64),
			password VARCHAR(64),
			status INT(4),
			createtime INT(10)
			);`
	ModifyDB(sql)
}

//创建文章表

func CreateTableWithArtcle() {
	sql := `CREATE TABLE IF NOT EXISTS article(
		id int(4) primary key auto_increment not null,
		title varchar(30),
		author varchar(20),
		tags varchar(30),
		short varchar(255),
		content longtext,
		createtime int(10)
		);`
	ModifyDB(sql)
}

//

func QueryRowDB(sql string) *sql.Row {
	return db.QueryRow(sql)
}

func MD5(text string) string {
	ctx := md5.New()
	ctx.Write([]byte(text))
	return hex.EncodeToString(ctx.Sum(nil))
}

//将传入的时间戳转为时间
func SwitchTimeStampToData(timestamp int64) string {
	t := time.Unix(timestamp, 0)
	return t.Format("2006-01-02 15:04:05")
}

func SwitchMarkdownToHtml(context string) template.HTML {
	markdown := blackfriday.MarkdownCommon([]byte(context))

	//获取到html
	doc, _ := goquery.NewDocumentFromReader(bytes.NewReader(markdown))

	/*
		对document进程查询 选择器和css的语法一样
		第一个参数： i是查询到的第几个元素
		第二个参数: selection就是查询到的元素
	*/
	doc.Find("code").Each(func(i int, selection *goquery.Selection) {
		light, _ := syntaxhighlight.AsHTML([]byte(selection.Text()))
		selection.SetHtml(string(light))
		fmt.Println(selection.Html())
		fmt.Println("light:", string(light))
		fmt.Println("\n\n\n")
	})

	htmlString, _ := doc.Html()
	return template.HTML(htmlString)
}
