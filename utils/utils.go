package utils

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"html/template"
	"time"

	"github.com/PuerkitoBio/goquery"

	"github.com/russross/blackfriday"
	"github.com/sourcegraph/syntaxhighlight"
)

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
