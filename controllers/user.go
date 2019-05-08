package controllers

import (
	"github.com/astaxie/beego"
	"novel/common"
	"novel/common/helper"
	"novel/models"
	"strconv"
	"strings"
	"time"
)

type UserController struct {
	BaseController
}

func (view *UserController) GetLogin() {
	view.Layout = "layouts/app.tpl"
	view.TplName = "auth/login.tpl"
}

func (c *UserController) PostLogin() {
	username := c.GetString("username")
	password := c.GetString("password")

	if !helper.Isset(username) {
		c.Error("用户名不能为空！")
		return
	}
	if !helper.Isset(password) {
		c.Error("密码不能为空！")
		return
	}

	user := models.GetUserByUserName(username)
	if user == nil {
		c.Error("用户不存在")
		return
	}
	if user.Status != 1 {
		c.Error("对不起您的账户已被冻结")
	}

	if user.Password != helper.EnUsernamePassword(user.Username, password) {
		c.Error("密码错误")
		return
	}

	c.SetSession(common.SESSION_USER_KEY, *user)
	c.Success("登录成功")
}

func (view *UserController) GetRegister() {
	flash := beego.ReadFromRequest(&view.Controller)

	if notice, ok := flash.Data["notice"]; ok {
		view.Data["Username"] = notice
	}
	if msg, ok := flash.Data["error"]; ok {
		message := strings.Split(msg,":")
		view.Data["ErrorField"] = message[0]
		view.Data["ErrorMessage"] = message[1]
	}

	view.Layout = "layouts/app.tpl"
	view.TplName = "auth/register.tpl"
}

func (c *UserController) PostRegister() {
	flash := beego.NewFlash()

	username := c.GetString("reg_username")
	password := c.GetString("reg_password")
	retPassword := c.GetString("password_confirmation")

	if !helper.Isset(username) {
		flash.Error("%s:用户名不能为空！", "username")
		flash.Store(&c.Controller)
		c.Redirect("/user/register",302)
		return
	}
	if !helper.Isset(password) {
		flash.Notice(username)
		flash.Error("%s:密码不能为空！", "password")
		flash.Store(&c.Controller)
		c.Redirect("/user/register",302)
		return
	}
	if password != retPassword {
		flash.Notice(username)
		flash.Error("%s:密码与确认密码不匹配！", "retPassword")
		flash.Store(&c.Controller)
		c.Redirect("/user/register",302)
		return
	}

	if models.GetUserByUserName(username) != nil {
		flash.Notice(username)
		flash.Error("%s:用户已存在！", "username")
		flash.Store(&c.Controller)
		c.Redirect("/user/register",302)
		return
	}

	var user models.User
	user.Username = username
	user.Password = helper.EnUsernamePassword(username, password)
	user.LastLoginIp = c.Ctx.Input.IP()
	loginDate,err := strconv.Atoi(strconv.FormatInt(time.Now().Unix(), 10))
	if err == nil {}
	user.RegisterDate = loginDate

	if _, err := models.AddUser(&user); err == nil {
		c.Redirect("/user/login",302)
	} else {
		flash.Notice(username)
		flash.Error("%s:注册失败！", "register")
		flash.Store(&c.Controller)
		c.Redirect("/user/register",302)
	}
}

func (c *UserController) Logout() {
	c.DelSession(common.SESSION_USER_KEY)

	if c.IsAjax() {
		c.Success("登出成功")
	}
	c.Redirect(c.Ctx.Request.Referer(), 302)
}