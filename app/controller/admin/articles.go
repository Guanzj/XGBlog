package admin

import (
	"XGBlog/app/loger"
	"XGBlog/app/protocol"
	"XGBlog/app/service/admin"
	"XGBlog/app/util"
	"XGBlog/app/validate"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Articles 文章的结构体
type Articles struct {
}

func (*Articles) getLogTitle() string {
	return "ctrller-admin-articles-"
}

// AddParams 新建文章结构体
type AddParams struct {
	Title       string `json:"Title" validate:"gt=4"`
	CateID      int    `json:"CateID" validate:"gt=0"`
	Description string `json:"Description"`
	Keywords    string `json:"Keywords"`
	Contents    string `json:"Contents" validate:"required"`
	ImgPath     string `json:"ImgPath"`
	PublishTime string `json:"PublishTime" validate:"required"`
	ShowType    int    `json:"ShowType" validate:"required"`
	Status      int    `json:"Status" validate:"required"`
}

//Add 添加文章
func (a *Articles) Add(c *gin.Context) {
	resp := &protocol.Resp{Ret: -1, Msg: "", Data: ""}
	var addParams AddParams
	err := c.ShouldBindJSON(&addParams)
	jsonstr, _ := json.Marshal(addParams)
	loger.Default().Info("Articles-Add-Params:", string(jsonstr))
	if err != nil {
		loger.Default().Info(a.getLogTitle(), "Add-error1:", err.Error())
		resp.Ret = -1
		resp.Msg = "参数错误"
		c.JSON(http.StatusOK, resp)
		return
	}

	validator, _ := validate.Default()
	if check := validator.CheckStruct(addParams); !check {
		resp.Ret = -1
		resp.Msg = validator.GetOneError()
		c.JSON(http.StatusOK, resp)
		return
	}

	if !util.IsTimeStr(addParams.PublishTime) {
		resp.Ret = -1
		resp.Msg = "发布时间格式错误"
		c.JSON(http.StatusOK, resp)
		return
	}
	if addParams.ShowType <= 0 {
		addParams.ShowType = 2
	}
	params := &admin.ArticleParams{}
	params.CateID = addParams.CateID
	params.Title = addParams.Title
	params.Description = addParams.Description
	params.Keywords = addParams.Keywords
	params.Contents = addParams.Contents
	params.ImgPath = addParams.ImgPath
	params.PublishTime = addParams.PublishTime
	params.ShowType = addParams.ShowType
	params.Status = addParams.Status
	adminID, _ := c.Get("adminID")
	username, _ := c.Get("username")
	fmt.Println(adminID)
	fmt.Println(username)
	// aopID, _ := strconv.Atoi(adminID.(string))
	// params.OpID = aopID
	params.OpID = 1
	// params.OpUser = username.(string)
	params.OpUser = "admin"
	articleServ := admin.Articles{}
	resp = articleServ.Add(params)
	c.JSON(http.StatusOK, resp)
}

//GetList 文章列表
func (a *Articles) GetList(c *gin.Context) {
	resp := &protocol.Resp{Ret: 0, Msg: "2", Data: ""}
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		page = 1
	}
	pageSize, err := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	if err != nil {
		pageSize = 10
	}
	cateID, err := strconv.Atoi(c.DefaultQuery("cateID", "0"))
	if err != nil {
		cateID = 0
	}
	articleServ := admin.Articles{}
	resp = articleServ.GetList(page, pageSize, cateID, []string{"articleID", "cateID", "title", "status", "description", "keywords", "imgPath", "opId", "opUser", "createTime", "modifyTime"})
	c.JSON(http.StatusOK, resp)
}

//Delete 删除文章
func (a *Articles) Delete(c *gin.Context) {
	resp := &protocol.Resp{Ret: 0, Msg: "", Data: ""}
	articleID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		loger.Default().Info(a.getLogTitle(), "Delete-error1:", err.Error())
		resp.Ret = -1
		resp.Msg = "参数错误"
		c.JSON(http.StatusOK, resp)
		return
	}
	articleServ := admin.Articles{}
	resp = articleServ.Delete(articleID)
	c.JSON(http.StatusOK, resp)
}

//Update 更新文章
func (a *Articles) Update(c *gin.Context) {
	resp := &protocol.Resp{Ret: -1, Msg: "", Data: ""}
	articleID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		loger.Default().Info(a.getLogTitle(), "Update-error1:", err.Error())
		resp.Ret = -1
		resp.Msg = "参数错误"
		c.JSON(http.StatusOK, resp)
		return
	}
	var addParams AddParams
	err = c.ShouldBindJSON(&addParams)
	jsonstr, _ := json.Marshal(addParams)
	loger.Default().Info("Articles-Update-Params:", string(jsonstr))
	if err != nil {
		loger.Default().Info(a.getLogTitle(), "Update-error2:", err.Error())
		resp.Ret = -1
		resp.Msg = "参数错误"
		c.JSON(http.StatusOK, resp)
		return
	}

	validator, _ := validate.Default()
	if check := validator.CheckStruct(addParams); !check {
		resp.Ret = -1
		resp.Msg = validator.GetOneError()
		c.JSON(http.StatusOK, resp)
		return
	}

	if !util.IsTimeStr(addParams.PublishTime) {
		resp.Ret = -1
		resp.Msg = "发布时间格式错误"
		c.JSON(http.StatusOK, resp)
		return
	}
	if addParams.ShowType <= 0 {
		addParams.ShowType = 2
	}
	params := &admin.ArticleParams{}
	params.CateID = addParams.CateID
	params.Title = addParams.Title
	params.Description = addParams.Description
	params.Keywords = addParams.Keywords
	params.Contents = addParams.Contents
	params.ImgPath = addParams.ImgPath
	params.PublishTime = addParams.PublishTime
	params.ShowType = addParams.ShowType
	params.Status = addParams.Status
	adminID, _ := c.Get("adminID")
	username, _ := c.Get("username")
	fmt.Println(adminID)
	fmt.Println(username)
	// aopID, _ := strconv.Atoi(adminID.(string))
	// params.OpID = aopID
	// params.OpUser = username.(string)

	params.OpID = 1
	params.OpUser = "admin"

	articleServ := admin.Articles{}
	resp = articleServ.Update(articleID, params)
	c.JSON(http.StatusOK, resp)
}

// Show 展示文章
func (a *Articles) Show(c *gin.Context) {
	resp := &protocol.Resp{Ret: 0, Msg: "", Data: ""}
	articleID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		loger.Default().Info(a.getLogTitle(), "Show-error1:", err.Error())
		resp.Ret = -1
		resp.Msg = "参数错误"
		c.JSON(http.StatusOK, resp)
		return
	}
	articleServ := admin.Articles{}
	resp = articleServ.Detail(articleID)
	c.JSON(http.StatusOK, resp)
}
