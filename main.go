package main

import (
	_ "beegotest/routers"
	utils "beegotest/utils"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

func main() {

	logs.SetLevel(beego.LevelInformational)
	logs.SetLogger("file", `{"filename": "logs/test.log"}`)
	utils.InitMysql()
	beego.Run()

}
