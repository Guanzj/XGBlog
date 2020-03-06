package front

import (
	"XGBlog/app/db/es"
	"XGBlog/app/db/mysql"
	"XGBlog/app/loger"
	"XGBlog/app/model"
	"XGBlog/app/protocol"
	"XGBlog/app/util"
	"context"
	"reflect"

	"github.com/olivere/elastic"
	// "context"
	// "github.com/olivere/elastic"
	// "myblog-api/app/db/es"
	// "myblog-api/app/loger"
)

// ArticleDetails 文章详情
type ArticleDetails struct {
	model.Articles        //在model文件夹
	Contents       string `json:"contents"`
	ShowType       int    `json:"show_type"`
}

// Articles 文章的结构体
type Articles struct {
}

// getLogTitle 获取 文章 日志的标题？？？？
func (*Articles) getLogTitle() string {
	return "service-admin-articles-"
}

// GetListFromMysql 分页获取文章列表mysql----id等信息有错误
func (a *Articles) GetListFromMysql(page int, pageSize int, cateID int, fields []string) (resp *protocol.Resp) {
	resp = &protocol.Resp{Ret: -1, Msg: "", Data: ""}
	db := mysql.Default().GetConn()
	defer db.Close()
	offset := (page - 1) * pageSize

	articles := make([]model.Articles, 0)

	if cateID > 0 {
		db = db.LogMode(true).Where("cateID = ?", cateID)
	}

	if err := db.Debug().LogMode(true).Select(fields).Offset(offset).Limit(pageSize).Order("articleID desc").Find(&articles).Error; err != nil {
		loger.Default().Error(a.getLogTitle(), "GetListForMysql-error1:", err.Error())
		resp.Msg = "系统错误"
		return resp
	}
	// fmt.Println(articles)
	resp.Ret = 0
	resp.Data = articles
	// fmt.Println(resp)
	return resp
}

// //GetListFromMysql 测试用
// func (a *Articles) GetListFromMysql(page int, pageSize int, cateID int, fields []string) (resp *protocol.Resp) {
// 	resp = &protocol.Resp{Ret: -1, Msg: "", Data: ""}
// 	db := mysql.Default().GetConn()
// 	defer db.Close()
// 	offset := (page - 1) * pageSize
// 	fmt.Println(offset)

// 	articles := make([]model.MbArticles, 0)

// 	db.Debug().LogMode(true).Select(fields).Find(&articles)

// 	fmt.Println(articles)
// 	resp.Ret = 0
// 	resp.Data = articles
// 	fmt.Println(resp)
// 	return resp
// }

// GetListFromEs 分页获取文章列表es
// Elastic是Go编程语言的Elasticsearch客户端
func (a *Articles) GetListFromEs(page int, pageSize int, cateID int, fields []string) (resp *protocol.Resp) {
	resp = &protocol.Resp{Ret: -1, Msg: "", Data: ""}

	esconn := es.Default().GetConn()
	ctx := context.Background()
	articles := make([]model.Articles, 0)

	query := esconn.Search().
		Index("myblog").
		Type("mb_articles").
		Size(pageSize).
		From((page-1)*pageSize).
		Sort("modify_time", false).
		Pretty(true)
	boolQuery := elastic.NewBoolQuery()
	searchQuery := boolQuery.Must(elastic.NewTermQuery("status", 1))
	if cateID > 0 {
		searchQuery = searchQuery.Filter(elastic.NewTermQuery("cateID", cateID))
	}
	query = query.Query(searchQuery)
	result, err := query.Do(ctx)
	if err != nil {
		loger.Default().Error(a.getLogTitle(), "GetListForEs-error1:", err.Error())
		resp.Msg = "系统错误"
		return resp
	}
	var typ model.Articles
	for _, item := range result.Each(reflect.TypeOf(typ)) { //从搜索结果中取数据的方法
		t := item.(model.Articles)
		t.ModifyTime = util.DateToDateTime(t.ModifyTime)
		t.CreateTime = util.DateToDateTime(t.CreateTime)
		articles = append(articles, t)

	}
	resp.Ret = 0
	resp.Data = articles
	return resp
}

//GetArticleDetail 获取文章详情---此处有问题，查不到id的显示未系统错误而不是文章不存在
func (a *Articles) GetArticleDetail(articleID int) (resp *protocol.Resp) {
	resp = &protocol.Resp{Ret: -1, Msg: "", Data: ""}
	articleContent := model.ArticlesContents{}
	articleDetails := ArticleDetails{}
	db := mysql.Default().GetConn()
	defer db.Close()
	if err := db.Where("articleID = ?", articleID).Where("status = ?", 1).First(&articleDetails).Error; err != nil {
		loger.Default().Error(a.getLogTitle(), "GetArticleDetail-error1:", err.Error())
		resp.Msg = "系统错误"
		return resp
	}
	if articleDetails.ArticleID <= 0 {
		resp.Msg = "文章不存在"
		return resp
	}
	if err := db.Where("articleID = ?", articleID).First(&articleContent).Error; err != nil {
		loger.Default().Error(a.getLogTitle(), "GetArticleDetail-error2:", err.Error())
		resp.Msg = "系统错误"
		return resp
	}
	resp.Ret = 0
	articleDetails.Contents = articleContent.Contents
	articleDetails.ShowType = articleContent.ShowType
	resp.Data = articleDetails
	return resp
}

//GetArticleCate 获取文章类型
func (a *Articles) GetArticleCate() (resp *protocol.Resp) {
	resp = &protocol.Resp{Ret: -1, Msg: "", Data: ""}
	articleCates := make([]model.ArticlesCate, 0)
	db := mysql.Default().GetConn()
	defer db.Close()
	if err := db.Debug().Order("orderby asc").Find(&articleCates).Error; err != nil {
		loger.Default().Error(a.getLogTitle(), "GetArticleCate-error1:", err.Error())
		resp.Msg = "系统错误"
		return resp
	}
	resp.Ret = 0
	resp.Data = articleCates
	return resp
}
