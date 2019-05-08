package models

import (
	"github.com/astaxie/beego/orm"
	"strings"
)

type Novel struct {
	Id              int    `orm:"column(id);auto"`
	Name            string `orm:"column(name);size(100)" description:"小说名称"`
	Author          string `orm:"column(author);size(50)" description:"作者"`
	Category        string `orm:"column(category);size(100)" description:"分类"`
	ContentValidity string `orm:"column(content_validity)" description:"内容简介"`
	Status          int8   `orm:"column(status)" description:"状态 1:连载中 2:完结"`
	LengthCont      int    `orm:"column(length_cont)" description:"小说长度"`
	CrawlId         int    `orm:"column(crawl_id)" description:"外部小说id"`
	CreatedAt       int    `orm:"column(created_at)" description:"创建时间"`
	UpdatedAt       int    `orm:"column(updated_at)" description:"更新时间"`
}

func (t *Novel) TableName() string {
	return "novel"
}

func init() {
	orm.RegisterModel(new(Novel))
}

func NovelFindAll(query map[string]string, limit int, sortby []string) [] Novel {

	o := orm.NewOrm()
	qs := o.QueryTable(new(Novel))

	//where:
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		} else {
			qs = qs.Filter(k, v)
		}
	}
	//Limit
	if limit >= 0 {
		qs = qs.Limit(limit)
	}
	// order by:
	if len(sortby) != 0 {
		qs = qs.OrderBy(sortby...)
	}

	var list []Novel
	if _, err := qs.All(&list); err == nil {
		return *&list
	}
	return nil
}

func NovelFindOne(query map[string]interface{}) *Novel {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Novel))

	//where:
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		} else {
			qs = qs.Filter(k, v)
		}
	}
	//result:
	var data Novel
	if err := qs.One(&data); err == nil {
		return &data
	}
	return nil
}

func NovelSearchByName(name string) *Novel {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Novel))

	cond := orm.NewCondition()
	condQuery := cond.And("name__icontains", name).Or("author__icontains", name)
	qs = qs.SetCond(condQuery)
	//result:
	var data Novel
	if err := qs.One(&data); err == nil {
		return &data
	}
	return nil
}

func GetNovelById(id int) (v *Novel) {
	o := orm.NewOrm()
	v = &Novel{Id: id}
	if err := o.Read(v); err == nil {
		return v
	}
	return nil
}
