package routers

import (
	"github.com/astaxie/beego"
	controllers "github.com/gxin0426/beegolearn/controllers"
)

func init() {
	//beego.Router("/", &controllers.MainController{})

	//beego.Router("/register")

	beego.Router("/", &controllers.HomeController{})
	beego.Router("/register", &controllers.RegisterController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/exit", &controllers.ExitController{})
	beego.Router("/article/add", &controllers.AddArticleController{})
}
