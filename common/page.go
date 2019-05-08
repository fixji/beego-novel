package common

import (
	"github.com/astaxie/beego"
	"math"
	"net/http"
	"net/url"
	"strconv"
)

//分页
type Page struct {
	CurrentPage int //获取当前页的页码
	Data       interface{}
	FirstPageUrl string //获取第一页 URL
	From int //计数
	LastPage int //获取最后一页的页码
	LastPageUrl string //获取最后一页的 URL
	NextPageUrl string //获取下一页的 URL
	Path string //URL
	PerPage int //每页显示的项目数
	PerPageUrl string //获取上一页的 URL
	To int //计数量
	Total int64 //判断存储器中匹配的所有项目总数
	PageUrls []string
}

func Paginate(currentPage int, data interface{}, req *http.Request, total int64) Page {
	appUrl := beego.AppConfig.String("APP_URL")
	query := req.URL.Query()
	query.Set("page", "1")
	firstPageUrl := req.URL.Path + "?" + query.Encode()
	from :=  currentPage * LIMIT + 1

	lastPage := int(math.Ceil(float64(total) / float64(LIMIT)))
	query.Set("page", strconv.Itoa(lastPage))
	lastPageUrl := req.URL.Path + "?" + query.Encode()

	nextPageUrl := ""
	if currentPage < lastPage {
		query.Set("page", strconv.Itoa(currentPage + 1))
		nextPageUrl = req.URL.Path + "?" + query.Encode()
	}

	perPageUrl := ""
	if currentPage < lastPage && currentPage > 1 {
		query.Set("page", strconv.Itoa(currentPage - 1))
		perPageUrl = req.URL.Path + "?" + query.Encode()
	}

	path := req.RequestURI

	pageUrls := PageUrls(req.URL.Path, req.URL.Query(), currentPage, lastPage)
	return Page{
		CurrentPage: currentPage,
		Data: data,
		FirstPageUrl: firstPageUrl,
		From: from,
		LastPage: lastPage,
		LastPageUrl: lastPageUrl,
		NextPageUrl: nextPageUrl,
		Path: appUrl + path,
		PerPage: LIMIT,
		PerPageUrl: perPageUrl,
		Total: total,
		PageUrls: pageUrls,
	}
}

func PageUrls(url string, query url.Values, currentPage int,  lastPage int) []string {
	var urls []string


	query.Set("page", strconv.Itoa(currentPage))
	currentPageUrl := "<a href='"+ url + "?" + query.Encode() +"'>"+ strconv.Itoa(currentPage) +"</a>"
	if lastPage < 7 {
		for i := 1; i < lastPage + 1 ; i++  {
			query.Set("page", strconv.Itoa(i))
			urls = append(urls, "<a href='"+ url + "?" + query.Encode() +"'>"+ strconv.Itoa(i) +"</a>")
		}
	} else {
		if currentPage < EACHSIDE + 2 {
			for i := 1; i < EACHSIDE + 4; i++  {
				query.Set("page", strconv.Itoa(i))
				urls = append(urls, "<a href='"+ url + "?" + query.Encode() +"'>"+ strconv.Itoa(i) +"</a>")
			}
			urls = append(urls, "...")
			query.Set("page", strconv.Itoa(lastPage))
			urls = append(urls, "<a href='"+ url + "?" + query.Encode() +"'>"+ strconv.Itoa(lastPage) +"</a>")
		} else if currentPage > EACHSIDE && currentPage < lastPage - EACHSIDE - 1 {
			query.Set("page", "1")
			urls = append(urls, "<a href='"+ url + "?" + query.Encode() +"'>1</a>")
			urls = append(urls, "...")

			for i := currentPage - EACHSIDE; i < currentPage + EACHSIDE + 1; i++ {
				query.Set("page", strconv.Itoa(i))
				urls = append(urls, "<a href='"+ url + "?" + query.Encode() +"'>"+ strconv.Itoa(i) +"</a>")
			}

			urls = append(urls, "...")
			query.Set("page", strconv.Itoa(lastPage))
			urls = append(urls, "<a href='"+ url + "?" + query.Encode() +"'>"+ strconv.Itoa(lastPage) +"</a>")
		} else {
			query.Set("page", "1")
			urls = append(urls, "<a href='"+ url + "?" + query.Encode() +"'>1</a>")
			urls = append(urls, "...")
			for i := lastPage - EACHSIDE - 2; i <= lastPage ; i++  {
				query.Set("page", strconv.Itoa(i))
				urls = append(urls, "<a href='"+ url + "?" + query.Encode() +"'>"+ strconv.Itoa(i) +"</a>")
			}
		}
	}

	for index := range urls {
		if currentPageUrl == urls[index] {
			urls[index] = "<a aria-current='page' class='current'><span>" + strconv.Itoa(currentPage) + "</span></a>"
		}
	}

	return urls
}