package common

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"strconv"
	"time"
)

//模版函数
func InitTempLate() {

	err := beego.AddFuncMap("date_time", dateTime)
	if err != nil {}

	err = beego.AddFuncMap("str_limit", strLimit)
	if err != nil {}
}

/**
由于mysql存的是int类型的时间戳 时间转化模版加层优化
*/
func dateTime(timestamp int, format string)(out string) {
	t, err := strconv.ParseInt(strconv.Itoa(timestamp), 10, 64)
	if err != nil {
		logs.Trace("template Execute err:", err)
	}
	tm := time.Unix(t, 0)
	return beego.Date(tm, format)
}

func strLimit(str string, limit int, end string) string {
	if len(str) > limit {
		str = beego.Substr(str, 0, limit)
	}
	return str + end
}