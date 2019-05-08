package helper

import (
	"crypto/md5"
	"encoding/hex"
	"reflect"
	"strings"
)

//辅助函数

func Isset(params interface{}) bool {
	//初始化变量
	var (
		flag          bool = false
		default_value reflect.Value
	)

	r := reflect.ValueOf(params)

	//获取对应类型默认值
	default_value = reflect.Zero(r.Type())
	//由于params 接口类型 所以default_value也要获取对应接口类型的值 如果获取不为接口类型 一直为返回false
	if !reflect.DeepEqual(r.Interface(), default_value.Interface()) {
		flag = true
	}
	return flag
}

func EnUsernamePassword(username string, str string) string {
	username = strings.Trim(username, " ")
	str = strings.Trim(str, " ")

	h := md5.New()
	h.Write([]byte(str))
	str = hex.EncodeToString(h.Sum(nil))
	h.Reset()
	h.Write([]byte(str + username))
	return hex.EncodeToString(h.Sum(nil))
}