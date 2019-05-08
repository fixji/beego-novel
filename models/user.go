package models

import (
	"github.com/astaxie/beego/orm"
)

type User struct {
	Id            int    `orm:"column(id);auto" description:"用户id"`
	Username      string `orm:"column(username);size(30)" description:"用户名"`
	Phone         string `orm:"column(phone);size(20)" description:"手机号"`
	Password      string `orm:"column(password);size(100)" description:"密码"`
	Status        int   `orm:"column(status)" description:"状态 2:已冻结 1:可用"`
	LastLoginIp   string `orm:"column(last_login_ip);size(50)" description:"最后登录ip"`
	RegisterDate  int    `orm:"column(register_date)" description:"注册时间"`
	LastLoginDate int    `orm:"column(last_login_date)" description:"最后登录时间"`
	LoginCount    int    `orm:"column(login_count)" description:"登录次数"`
}

func (t *User) TableName() string {
	return "user"
}

func init() {
	orm.RegisterModel(new(User))
}

// AddUser insert a new User into database and returns
// last inserted Id on success.
func AddUser(m *User) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetUserById retrieves User by Id. Returns error if
// Id doesn't exist
func GetUserByUserName(username string) *User {
	o := orm.NewOrm()
	qs := o.QueryTable(new(User))
	qs = qs.Filter("username", username)

	var data User
	if err := qs.One(&data); err == nil {
		return &data
	}
	return nil
}