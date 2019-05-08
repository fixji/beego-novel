package controllers

import (
	"novel/common"
	"novel/common/helper"
	"novel/models"
	"strconv"
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
	view.Layout = "layouts/app.tpl"
	view.TplName = "auth/register.tpl"
}

func (c *UserController) PostRegister() {
	username := c.GetString("reg_username")
	password := c.GetString("reg_password")
	retPassword := c.GetString("password_confirmation")

	if !helper.Isset(username) {
		c.Error("用户名不能为空！")
		return
	}
	if !helper.Isset(password) {
		c.Error("密码不能为空！")
		return
	}
	if password != retPassword {
		c.Error("密码与确认密码不匹配！")
		return
	}

	if models.GetUserByUserName(username) != nil {
		c.Error("用户已存在")
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
		c.Success("注册成功")
	} else {
		c.Error(err.Error())
	}
}

func (c *UserController) Logout() {
	c.DelSession(common.SESSION_USER_KEY)

	if c.IsAjax() {
		c.Success("登出成功")
	}
	c.Redirect(c.Ctx.Request.Referer(), 302)
}