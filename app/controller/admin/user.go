/*
 * @Author: moonmist.guan
 * @Date: 2020-03-05 15:43:15
 * @LastEditTime: 2020-03-06 17:55:06
 * @FilePath: /XGBlog/app/controller/admin/user.go
 * @Description:
 */

package admin

import (
	"XGBlog/app/protocol"
	"net/http"

	"github.com/gin-gonic/gin"
)

// User 后台用户
type User struct {
}

//Info 用户信息
type Info struct {
	Roles        []string `json:"roles"`
	Introduction string   `json:"introduction"`
	Avatar       string   `json:"avatar"`
	Name         string   `json:"name"`
}

//Show 用户信息
func (User) Show(c *gin.Context) {
	resp := &protocol.Resp{Ret: 0, Msg: "", Data: ""}
	roles := []string{"admin"}
	resp.Data = &Info{
		Roles:        roles,
		Introduction: "我是管理员我怕谁",
		Avatar:       "https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif",
		Name:         "admin",
	}
	c.JSON(http.StatusOK, resp)
}
