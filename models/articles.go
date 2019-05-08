package models

import (
	"github.com/astaxie/beego/orm"
	"strings"
)

type Articles struct {
	Id        int    `orm:"column(id);auto"`
	NovelId   int    `orm:"column(novel_id)" description:"小说id"`
	Title     string `orm:"column(title);size(100)" description:"小说文章标题"`
	Sort      int    `orm:"column(sort)" description:"排序"`
	CreatedAt int    `orm:"column(created_at)" description:"创建时间"`
}

func (t *Articles) TableName() string {
	return "articles"
}

func init() {
	orm.RegisterModel(new(Articles))
}

func GetArticlesById(id int) (v *Articles) {
	o := orm.NewOrm()
	v = &Articles{Id: id}
	if err := o.Read(v); err == nil {
		return v
	}
	return nil
}


func ArticlesFindOne(query map[string]interface{}, sortby []string) *Articles {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Articles))

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
	// order by:
	if len(sortby) != 0 {
		qs = qs.OrderBy(sortby...)
	} else {
		return nil
	}
	//result:
	var data Articles
	if err := qs.One(&data); err == nil {
		return &data
	}
	return nil
}

func ArticlesFindAll(query map[string]string, offset int, limit int, sortby []string) [] Articles {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Articles))

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
	//Limit:
	if limit >= 0 {
		qs = qs.Limit(limit)
	}
	//Offset
	if offset >= 0 {
		qs = qs.Offset(offset)
	}
	// order by:
	if len(sortby) != 0 {
		qs = qs.OrderBy(sortby...)
	}
	//result:
	var list []Articles
	if _, err := qs.All(&list); err == nil {
		return *&list
	}
	return nil
}

func ArticlesCount(query map[string]string) int64 {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Articles))

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
	count, err := qs.Count()
	if err != nil{
		return 0
	}
	return count
}
