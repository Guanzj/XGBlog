/*
 * @Author: moonmist.guan
 * @Date: 2020-03-05 14:59:33
 * @LastEditTime: 2020-03-06 23:43:14
 * @FilePath: /XGBlog/app/middleware/access.go
 * @Description:
 */

package middleware

import (
	"XGBlog/app/loger"
	"XGBlog/app/service/front"
	"encoding/json"
	"time"

	"github.com/gin-gonic/gin"
)

// AddAccessLog 添加路由通过的日志，返回一个gin的handle
func AddAccessLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		var accessLogServ front.AccessLog
		params := &front.AccessLogParams{}
		params.IP = c.ClientIP()
		params.Timestamp = time.Now().Unix()
		params.Path = c.Request.Method + "|" + c.Request.URL.Path
		params.Date = time.Now().Format("2006-01-02")
		err := accessLogServ.Add(params)
		paramsstr, _ := json.Marshal(params)
		loger.Default().Info("AddAccessLog-params:", string(paramsstr))
		if err != nil {
			loger.Default().Error("AddAccessLog-error:", err.Error())
		}
		c.Next()
	}

}
