package controllers

import (
	"github.com/astaxie/beego"
	"html/template"
	"novel/common"
	"novel/models"
	"reflect"
)

type BaseController struct {
	beego.Controller
}

//这个函数主要是为了用户扩展用的，这个函数会在下面定义的这些 Method 方法之前执行，用户可以重写这个函数实现类似用户验证之类。
func (c *BaseController) Prepare() {
	if user, err := c.GetSession(common.SESSION_USER_KEY).(models.User); err {
		c.Data["User"] = user
	}
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.Data["xsrf_token"]= c.XSRFToken()
}

func (c *BaseController) Response(content interface{}, status int) {
	var result map[string]interface{}

	if reflect.TypeOf(content).String() == "string" {
		result = map[string]interface{} {
			"code": status,
			"message": content,
		}
	} else {
		result = map[string]interface{} {
			"code": status,
			"message": "success",
			"data" : content,
		}
	}

	c.Data["json"] = result
	c.ServeJSON()
	c.StopRun()
}

func (c *BaseController) Success(content interface{}) {

	var result map[string]interface{}

	if reflect.TypeOf(content).String() == "string" {
		result = map[string]interface{} {
			"code": 200,
			"message": content,
		}
	} else {
		result = map[string]interface{} {
			"code": 200,
			"message": "success",
			"data" : content,
		}
	}

	c.Data["json"] = result
	c.ServeJSON()
	c.StopRun()
}

func (c *BaseController) NotFund(message string) {
	result := map[string]interface{} {
		"code": 404,
		"message": message,
	}
	c.Data["json"] = result
	c.ServeJSON()
	c.StopRun()
}

func (c *BaseController) Error(message string) {
	result := map[string]interface{} {
		"code": 400,
		"message": message,
	}
	c.Data["json"] = result
	c.ServeJSON()
	c.StopRun()
}