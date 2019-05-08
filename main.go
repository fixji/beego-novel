package main

import (
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"novel/common"
	_ "novel/routers"
)

func main() {

	//初始化配置
	common.InitApp()
	//增加模版方法
	common.InitTempLate()

	beego.Run()
}

