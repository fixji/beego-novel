package routers

import (
	"github.com/astaxie/beego"
	"novel/controllers"
)

func init() {
	beego.ErrorController(&controllers.ErrorController{})

    beego.Router("/", &controllers.NovelController{}, "get:Index")
    beego.Router("/novel/:id", &controllers.NovelController{}, "get:Description")
    beego.Router("/novel/content/:id", &controllers.NovelController{}, "get:Content")
    beego.Router("/novel/search", &controllers.NovelController{}, "get:Search")


    //beego.Router("/api", &controllers.UserController{}, "post:Search")
    //尝试不同路由方法
	beego.Router("/user/login", &controllers.UserController{}, "get:GetLogin;post:PostLogin")
	beego.Router("/user/register", &controllers.UserController{}, "get:GetRegister;post:PostRegister")
	beego.Router("/user/logout", &controllers.UserController{}, "post:Logout")

}
