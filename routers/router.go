package routers

import (
	"github.com/astaxie/beego"
	controllers "github.com/gxin0426/beegolearn/controllers"
)

func init() {
	//beego.Router("/", &controllers.MainController{})

	//beego.Router("/register")

	beego.Router("/register", &controllers.RegisterController{})
}
