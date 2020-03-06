/*
 * @Author: moonmist.guan
 * @Date: 2020-03-05 15:43:15
 * @LastEditTime: 2020-03-06 15:22:48
 * @FilePath: /XGBlog/app/controller/admin/login.go
 * @Description:
 */

package admin

import (
	"XGBlog/app/loger"
	"XGBlog/app/protocol"
	"XGBlog/app/service/admin"
	"XGBlog/app/validate"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Login 后台登录
type Login struct {
}

//LoginParams 登录参数
type LoginParams struct {
	Username string `json:"username" form:"username"  binding:"required"`
	Password string `json:"password" form:"password"  binding:"required"`
	Code     string `json:"code" `
}

// type LoginParams struct {
// 	Username string `json:"username" validate:"gt=4" form:"username"`
// 	Password string `json:"password" validate:"gt=6"`
// 	Code     string `json:"code" validate:"len=6"`
// }

func (*Login) getLogTitle() string {
	return "ctrller-admin-login-"
}

//Login 登录
func (login *Login) Login(c *gin.Context) {
	resp := protocol.Resp{Ret: 0, Msg: "", Data: ""}
	var loginParams LoginParams
	// 前端使用form提交时由下面的方法读取参数
	// fmt.Println(c.PostForm("username"))
	// fmt.Println(c.PostForm("password"))
	err := c.ShouldBindJSON(&loginParams)
	fmt.Println(loginParams)
	if err != nil {
		loger.Default().Error(login.getLogTitle(), "Login-error1:", err.Error())
		resp.Ret = -1
		resp.Msg = "用户名或密码不能为空1"
		c.JSON(http.StatusOK, resp)
		return
	}
	username := loginParams.Username
	password := loginParams.Password
	code, err := strconv.Atoi(loginParams.Code)
	if err != nil {
		// loger.Default().Error(login.getLogTitle(), "Login-error2:", err.Error())
		// resp.Ret = -1
		// resp.Msg = "谷歌验证码错误"
		// c.JSON(http.StatusOK, resp)
		// return
	}
	fmt.Println("密码校验开始")

	adminServ := admin.Admins{}
	validator, _ := validate.Default()
	if check := validator.CheckStruct(loginParams); !check {
		resp.Ret = -1
		resp.Msg = validator.GetOneError()
		c.JSON(http.StatusOK, resp)
		return
	}

	resp = adminServ.Login(username, password, uint32(code))
	c.JSON(http.StatusOK, resp)
}
