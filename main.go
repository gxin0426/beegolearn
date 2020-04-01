package main

import (
	_ "github.com/gxin0426/beegolearn/routers"
	"github.com/gxin0426/beegolearn/utils"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

func main() {

	logs.SetLevel(beego.LevelInformational)
	logs.SetLogger("file", `{"filename": "logs/test.log"}`)
	utils.InitMysql()
	beego.Run()

}
