package controllers

import (
	"novel/common"
	"novel/common/helper"
	"novel/models"
	"strconv"
)

type NovelController struct {
	BaseController
}

func (view *NovelController) Index() {
	view.Data["Hot"] = models.NovelFindAll(nil,3, nil)

	sortBy := []string{"-updated_at"}
	view.Data["XianXia"] = models.NovelFindAll(map[string]string {"category": "武侠仙侠"},7, sortBy)
	view.Data["XuanHuan"] = models.NovelFindAll(map[string]string {"category": "玄幻奇幻"},7, sortBy)
	view.Data["DuShi"] = models.NovelFindAll(map[string]string {"category": "都市言情"},7, sortBy)

	view.Layout = "layouts/app.tpl"
	view.TplName = "novel/index.tpl"
}

func (view *NovelController) Description() {

	id, err := strconv.Atoi(view.Ctx.Input.Param(":id"))
	if !helper.Isset(id) {
		view.Data["ErrorMessage"] = "未找到该小说！"
		view.Abort("404")
	}
	if err != nil{}
	info := models.GetNovelById(id)
	if !helper.Isset(info) {
		view.Data["ErrorMessage"] = "未找到该小说！"
		view.Abort("404")
	}

	redis := common.NewRedis()
	click, err := redis.Incr("novel:" + strconv.Itoa(info.Id)).Result()
	if err != nil{}
	view.Data["Click"] = click
	redis.Close()

	view.Data["Info"] = info
	view.Data["Name"] = info.Name

	filter := map[string]interface{} {"novel_id": info.Id}
	view.Data["New"] = models.ArticlesFindOne(filter, []string{"-sort"})
	view.Data["First"] = models.ArticlesFindOne(filter, []string{"sort"})

	offset := 0
	page, err := strconv.Atoi(view.GetString("page"))
	if err != nil{}
	var sortBy []string
	sort := view.GetString("sort")
	if sort == "desc" {
		sortBy = append(sortBy, "-sort")
	} else {
		sortBy = append(sortBy, "sort")
	}
	if helper.Isset(page) && page > 0 {
		offset = (page - 1) * common.LIMIT
	} else {
		page = 1
	}
	data := models.ArticlesFindAll(map[string]string {"novel_id":strconv.Itoa(info.Id)}, offset, common.LIMIT, sortBy)
	count := models.ArticlesCount(map[string]string {"novel_id":strconv.Itoa(info.Id)})
	view.Data["Page"] = common.Paginate(page, data, view.Ctx.Request, count)
	view.Data["Sort"] = sort
	view.Layout = "layouts/app.tpl"
	view.TplName = "novel/desc.tpl"
}

func (view *NovelController) Content() {
	id, err := strconv.Atoi(view.Ctx.Input.Param(":id"))
	if err != nil{}
	article := models.GetArticlesById(id)
	if !helper.Isset(article) {
		view.Data["ErrorMessage"] = "章节丢失！"
		view.Abort("404")
	}
	view.Data["Article"] = article

	prevFilter := map[string]interface{} {"novel_id":strconv.Itoa(article.NovelId), "sort__lt": article.Sort}
	view.Data["PrevCont"] = models.ArticlesFindOne(prevFilter, []string{"-sort"})

	nextFilter := map[string]interface{} {"novel_id":strconv.Itoa(article.NovelId), "sort__gt": article.Sort}
	view.Data["NextCont"] = models.ArticlesFindOne(nextFilter, []string{"sort"})

	content := models.GetContentById(article.Id)
	if !helper.Isset(article) {
		view.Data["ErrorMessage"] = "章节丢失！"
		view.Abort("404")
	}
	view.Data["Cont"] = content.Content

	novel := models.GetNovelById(article.NovelId)
	view.Data["Novel"] = novel
	view.Data["Name"] = novel.Name
	view.Data["Title"] = article.Title
	view.Layout = "layouts/app.tpl"
	view.TplName = "novel/cont.tpl"
}

func (view *NovelController) Search() {
	name := view.GetString("name")

	if !helper.Isset(name) {
		view.Data["ErrorMessage"] = "未找到该小说！"
		view.Abort("404")
	}
	data := models.NovelSearchByName(name)
	if !helper.Isset(data) {
		view.Data["ErrorMessage"] = "未找到该小说！"
		view.Abort("404")
	}
	view.Redirect("/novel/" + strconv.Itoa(data.Id), 302)
}