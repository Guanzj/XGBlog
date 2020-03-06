package admin

import (
	"XGBlog/app/db/mysql"
	"XGBlog/app/loger"
	"XGBlog/app/model"
	"XGBlog/app/protocol"
	"strconv"
	"time"
)

// ArticleParams 文章结构体
type ArticleParams struct {
	CateID      int    `json:"cateID"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Keywords    string `json:"keywords"`
	ImgPath     string `json:"img_path"`
	OpID        int    `json:"op_id"`
	OpUser      string `json:"op_user"`
	Contents    string `json:"contents"`
	ShowType    int    `json:"show_type"`
	PublishTime string `json:"publish_time"`
	Status      int    `json:"status"`
}

// Detail 文章详情
type Detail struct {
	ID          int    `json:"id"`
	CateID      int    `json:"cateID"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Keywords    string `json:"keywords"`
	ImgPath     string `json:"img_path"`
	Contents    string `json:"contents"`
	ShowType    int    `json:"show_type"`
	Status      int    `json:"status"`
	PublishTime string `json:"publish_time"`
}

// Articles 文章集合
type Articles struct {
}

func (*Articles) getLogTitle() string {
	return "service-admin-articles-"
}

// ArticleList 文章列表
type ArticleList struct {
	Total    int              `json:"total"`
	Datalist []model.Articles `json:"datalist"`
}

//Add 添加文章
func (a *Articles) Add(params *ArticleParams) (resp *protocol.Resp) {
	resp = &protocol.Resp{Ret: -1, Msg: "", Data: ""}
	articles := model.Articles{
		CateID:      params.CateID,
		Title:       params.Title,
		Description: params.Description,
		Keywords:    params.Keywords,
		ImgPath:     params.ImgPath,
		OpID:        params.OpID,
		OpUser:      params.OpUser,
		ModifyTime:  params.PublishTime,
		Status:      params.Status,
		CreateTime:  time.Now().Format("2006-01-02 15:04:05"),
	}

	articlesContents := model.ArticlesContents{
		ShowType: params.ShowType,
		Contents: params.Contents,
	}
	//if articlesContents.GetShowTypeName() == "" {
	//	resp.Msg = "文章内容显示类型错误"
	//	return resp
	//}
	//添加articlesContents
	db := mysql.Default().GetConn()
	defer db.Close()
	// 开始事务
	tx := db.Begin()
	//添加articles
	err := db.Model(model.Articles{}).Create(&articles).Error
	if err != nil {
		loger.Default().Error(a.getLogTitle(), "Add-error1:", err.Error())
		resp.Msg = "系统错误"
		tx.Rollback()
		return resp
	}
	//获取插入记录的Id
	var articleID []int
	db.Raw("select LAST_INSERT_ID() as id").Pluck("articleID", &articleID)
	articlesContents.ArticleID = articleID[0]
	err = db.Create(&articlesContents).Error
	if err != nil {
		loger.Default().Error(a.getLogTitle(), "Add-error2:", err.Error())
		resp.Msg = "系统错误"
		tx.Rollback()
		return resp
	}
	//提交事务
	tx.Commit()
	resp.Ret = 0
	resp.Data = articlesContents
	return resp
}

//Update 更新文章
func (a *Articles) Update(id int, params *ArticleParams) (resp *protocol.Resp) {
	resp = &protocol.Resp{Ret: -1, Msg: "", Data: ""}
	//查询ID是否存在
	db := mysql.Default().GetConn()
	defer db.Close()
	count := 0
	if err := db.Model(model.Articles{}).Where("articleID = ?", id).Count(&count).Error; err != nil {
		loger.Default().Error(a.getLogTitle(), "Update-error1:", err.Error())
		resp.Msg = "系统错误"
		return resp
	}

	if count <= 0 {
		resp.Msg = "文章不存在"
		return resp
	}
	articles := model.Articles{
		CateID:      params.CateID,
		Title:       params.Title,
		Status:      params.Status,
		Description: params.Description,
		Keywords:    params.Keywords,
		ImgPath:     params.ImgPath,
		OpID:        params.OpID,
		OpUser:      params.OpUser,
		ModifyTime:  params.PublishTime,
	}

	articlesContents := model.ArticlesContents{
		ShowType: params.ShowType,
		Contents: params.Contents,
	}
	//if articlesContents.GetShowTypeName() == "" {
	//	resp.Msg = "文章内容显示类型错误"
	//	return resp
	//}
	// 开始事务
	tx := db.Begin()
	//添加articles
	err := db.Model(model.Articles{}).Where("articleID = ?", id).Update(&articles).Error
	if err != nil {
		loger.Default().Error(a.getLogTitle(), "Update-error2:", err.Error())
		resp.Msg = "系统错误"
		tx.Rollback()
		return resp
	}
	//获取插入记录的Id
	err = db.Model(model.ArticlesContents{}).Where("articleID = ?", id).Updates(&articlesContents).Error
	if err != nil {
		loger.Default().Error(a.getLogTitle(), "Update-error3:", err.Error())
		resp.Msg = "系统错误"
		tx.Rollback()
		return resp
	}
	//提交事务
	tx.Commit()
	resp.Ret = 0
	resp.Data = articles
	resp.Msg = "修改成功"
	return resp
}

//GetList 分页获取文章列表
func (a *Articles) GetList(page int, pageSize int, cateID int, fields []string) (resp *protocol.Resp) {
	resp = &protocol.Resp{Ret: -1, Msg: "1", Data: ""}
	db := mysql.Default().GetConn()
	defer db.Close()
	offset := (page - 1) * pageSize
	articleList := &ArticleList{}
	articles := make([]model.Articles, 0)
	total := 0
	if cateID > 0 {
		db = db.Where("cateID = ?", cateID)
	}
	db.Model(&model.Articles{}).Count(&total)
	if err := db.Select(fields).Offset(offset).Limit(pageSize).Order("articleID desc").Find(&articles).Error; err != nil {
		loger.Default().Error(a.getLogTitle(), "GetList-error1:", err.Error())
		resp.Msg = "系统错误"
		return resp
	}
	articleList.Datalist = articles
	articleList.Total = total
	resp.Ret = 0
	resp.Data = articleList
	return resp
}

//Delete 删除文章
func (a *Articles) Delete(id int) (resp *protocol.Resp) {
	resp = &protocol.Resp{Ret: -1, Msg: "", Data: ""}
	db := mysql.Default().GetConn()
	defer db.Close()
	if err := db.Where("articleID = ?", id).Delete(&model.Articles{}).Error; err != nil {
		loger.Default().Error(a.getLogTitle(), "Delete-error:", err.Error())
		resp.Msg = "系统错误"
		return resp
	}
	resp.Ret = 0
	resp.Msg = strconv.Itoa(id) + "删除成功"
	resp.Data = 200
	return resp
}

//Detail 文章详情
func (a *Articles) Detail(id int) (resp *protocol.Resp) {
	resp = &protocol.Resp{Ret: -1, Msg: "", Data: ""}
	db := mysql.Default().GetConn()
	defer db.Close()
	article := &model.Articles{}
	articleContent := &model.ArticlesContents{}

	if err := db.Where("articleID = ?", id).Find(article).Error; err != nil {
		loger.Default().Error(a.getLogTitle(), "Detail-error1:", err.Error())
		resp.Msg = "系统错误"
		return resp
	}
	if err := db.Where("articleID = ?", id).Find(articleContent).Error; err != nil {
		loger.Default().Error(a.getLogTitle(), "Detail-error2:", err.Error())
		resp.Msg = "系统错误"
		return resp
	}
	detail := &Detail{}
	detail.Title = article.Title
	detail.ID = id
	detail.CateID = article.CateID
	detail.Description = article.Description
	detail.Keywords = article.Keywords
	detail.ImgPath = article.ImgPath
	detail.Status = article.Status
	detail.PublishTime = article.ModifyTime
	detail.Contents = articleContent.Contents
	detail.ShowType = articleContent.ShowType
	resp.Data = detail
	resp.Ret = 0
	return resp
}
