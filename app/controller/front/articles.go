/*
 * @Author: moonmist.guan
 * @Date: 2020-03-05 15:37:25
 * @LastEditTime: 2020-03-06 01:49:02
 * @FilePath: /XGBlog/app/controller/front/articles.go
 * @Description:
 */

package front

import (
	"XGBlog/app/protocol"
	"XGBlog/app/service/front"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Articles 文章的结构体
type Articles struct {
}

// GetList 文章列表
func (Articles) GetList(c *gin.Context) {
	resp := &protocol.Resp{Ret: 0, Msg: "", Data: ""}
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		page = 1
	}
	pageSize, err := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	if err != nil {
		pageSize = 10
	}
	cateID, err := strconv.Atoi(c.DefaultQuery("cateID", "1"))
	if err != nil {
		cateID = 0
	}
	articleServ := front.Articles{}
	// 从es中获取列表
	// resp = articleServ.GetListFromEs(page, pageSize, cateID, []string{"articleID", "cateID", "title", "description", "modifyTime"})

	resp = articleServ.GetListFromMysql(page, pageSize, cateID, []string{"articleID", "cateID", "title", "status", "description", "keywords", "imgPath", "opId", "opUser", "createTime", "modifyTime"})
	c.JSON(http.StatusOK, resp)
}

//Show 文章详情
func (Articles) Show(c *gin.Context) {
	resp := &protocol.Resp{Ret: -1, Msg: "", Data: ""}
	id, err := strconv.Atoi(c.Param("id"))
	articleServ := front.Articles{}
	if err != nil || id <= 0 {
		resp.Ret = -1
		resp.Msg = "参数错误"
		c.JSON(http.StatusOK, resp)
		return
	}
	resp = articleServ.GetArticleDetail(id)
	c.JSON(http.StatusOK, resp)
}

//GetCategories 文章类型
func (Articles) GetCategories(c *gin.Context) {
	resp := &protocol.Resp{Ret: 0, Msg: "", Data: ""}
	articleServ := front.Articles{}
	resp = articleServ.GetArticleCate()
	c.JSON(http.StatusOK, resp)
}
