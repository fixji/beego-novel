package controllers

import (
	"github.com/astaxie/beego"
)

type ErrorController struct {
	beego.Controller
}

func (view *ErrorController) Error404() {
	view.Data["Message"] = view.Ctx.ResponseWriter.Header().Get("ErrorMessage")
	view.TplName = "errors/404.tpl"
}
