package common

import (
	"encoding/gob"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/go-redis/redis"
	_ "github.com/astaxie/beego/session/redis"
	"novel/models"
	"strconv"
)

//全局redis
func InitApp() {
	//初始化Mysql
	err := orm.RegisterDriver("mysql", orm.DRMySQL)
	if err != nil {}
	err = orm.RegisterDataBase("default", "mysql", beego.AppConfig.String("DB_DSN"))
	if err != nil {}
	orm.SetMaxIdleConns("default",10)
	orm.SetMaxOpenConns("default",15)
	//初始化Session
	InitSession()

}

func InitSession() {
	gob.Register(models.User{})
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.BConfig.WebConfig.Session.SessionName = "session_id"
	beego.BConfig.WebConfig.Session.SessionProvider = "redis"
	//表示链接的地址，连接池，访问密码,端口
	beego.BConfig.WebConfig.Session.SessionProviderConfig = beego.AppConfig.String("SESSION_DSN")
}

func NewRedis() *redis.Client {
	redisDb, err := strconv.Atoi(beego.AppConfig.String("REDIS_DB"))
	if err != nil {}
	return redis.NewClient(&redis.Options{
		Addr: beego.AppConfig.String("REDIS_HOST"),
		Password: beego.AppConfig.String("REDIS_PASSWORD"),
		DB: redisDb,
	})
}
