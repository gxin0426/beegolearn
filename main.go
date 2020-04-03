package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	_ "github.com/gxin0426/beegolearn/routers"
	utils "github.com/gxin0426/beegolearn/utils"
)

func main() {

	logs.SetLevel(beego.LevelInformational)
	logs.SetLogger("file", `{"filename": "logs/test.log"}`)
	utils.InitMysql()
	beego.Run()

}
