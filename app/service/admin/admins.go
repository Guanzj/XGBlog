/*
 * @Author: moonmist.guan
 * @Date: 2020-03-05 16:06:25
 * @LastEditTime: 2020-03-06 16:08:09
 * @FilePath: /XGBlog/app/service/admin/admins.go
 * @Description:
 */

package admin

import (
	"XGBlog/app/config"
	"XGBlog/app/db/mysql"
	"XGBlog/app/loger"
	"XGBlog/app/model"
	"XGBlog/app/protocol"
	"XGBlog/app/util"
	"fmt"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// Admins 后台管理
type Admins struct {
}

func (*Admins) getLogTitle() string {
	return "service-admin-login-"
}

//UserInfo 用户信息
type UserInfo struct {
	AdminID  int32  `json:"admin_id"`
	UserName string `json:"user_name"`
	Token    string `json:"token"`
}

//Login 登录---密码md5后有问题，
func (a *Admins) Login(username string, password string, code uint32) (resp protocol.Resp) {
	resp = protocol.Resp{Ret: -1, Msg: "", Data: ""}

	//校验谷歌验证码---暂时不用
	// gaCode, err := util.MkGaCode(config.Configs.GaSecret)
	// if err != nil {
	// 	loger.Default().Error(a.getLogTitle(), "Login-error1:", err.Error())
	// 	resp.Msg = "系统错误"
	// 	return resp
	// }

	// if code != gaCode {
	// 	resp.Msg = "谷歌验证码错误"
	// 	return resp
	// }

	db := mysql.Default().GetConn()
	defer db.Close()
	//查询用户
	admin := model.Admins{}
	db.Where("username=?", username).First(&admin)
	if util.IsEmpty(admin) {
		resp.Msg = "账号不存在"
		return resp
	}
	fmt.Printf("%T,%v\n", util.MkMd5(password), util.MkMd5(password))

	fmt.Printf("%T,%v\n", admin.Password, admin.Password)
	fmt.Println("客户端密码:", util.MkMd5(password))
	fmt.Println("服务端密码:", admin.Password)
	fmt.Println(util.MkMd5(password) == admin.Password)
	fmt.Println(strings.Compare(util.MkMd5(password), admin.Password))
	fmt.Println(strings.EqualFold(util.MkMd5(password), admin.Password))

	//检测密码是否正确
	// if util.MkMd5(password) != admin.Password {
	// 	loger.Default().Info(a.getLogTitle(), "admin:", admin)
	// 	loger.Default().Info(a.getLogTitle(), "服务端密码:", admin.Password)
	// 	//客户端使用了md5加密，所以服务端也应该存储md5后的密码
	// 	loger.Default().Info(a.getLogTitle(), "客户端密码:", util.MkMd5(password))
	// 	resp.Msg = "密码错误"
	// 	return resp
	// }
	if password != admin.Password {
		loger.Default().Info(a.getLogTitle(), "admin:", admin)
		loger.Default().Info(a.getLogTitle(), "服务端密码:", admin.Password)
		//客户端使用了md5加密，所以服务端也应该存储md5后的密码
		loger.Default().Info(a.getLogTitle(), "客户端密码:", util.MkMd5(password))
		resp.Msg = "密码错误"
		return resp
	}

	//生成token--注意过期时间--调用了config
	token, err := util.JwtEncode(jwt.MapClaims{"admin_id": fmt.Sprintf("%d", admin.AdminID), "username": admin.Username, "expr_time": fmt.Sprintf("%d", time.Now().Unix())}, []byte(config.Configs.JwtSecret))
	if err != nil {
		loger.Default().Error(a.getLogTitle(), "Login-error2:", err.Error())
		resp.Ret = -999
		resp.Msg = "系统错误"
		return resp
	}
	userInfo := UserInfo{
		AdminID:  admin.AdminID,
		UserName: admin.Username,
		Token:    token,
	}
	resp.Data = userInfo
	resp.Ret = 0
	return resp
}
