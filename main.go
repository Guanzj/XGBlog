/*
 * @Author: moonmist.guan
 * @Date: 2020-03-05 14:21:52
 * @LastEditTime: 2020-03-06 20:43:23
 * @FilePath: /XGBlog/main.go
 * @Description:
 */

package main

import (
	"XGBlog/app/config"
	"XGBlog/app/router"
	"XGBlog/app/util"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

// func main() {
// 	r := gin.Default()
// 	r.GET("/ping", func(c *gin.Context) {
// 		fmt.Println("hello world!")
// 		c.JSON(200, gin.H{
// 			"message": "pong",
// 		})
// 	})
// 	r.Run() // listen and serve on 0.0.0.0:8080
// }

func init() {
	// 获取当前项目所在路径
	appDir := util.GetAppDir()
	fmt.Println(appDir + "/config.ini")
	// 加载config文件
	config.Default(appDir + "/config.ini")

}

func main() {
	// 判断是生产还是开发环境
	if config.Configs.Env == "prd" {
		gin.SetMode(gin.ReleaseMode)
	}
	gin.SetMode(gin.DebugMode)
	// 启动路由器
	r := router.Default()
	r.Run()
	//启动到固定端口
	err := r.GetEngin().Run("0.0.0.0:" + config.Configs.HttpListenPort)
	if err != nil {
		fmt.Println("start service error!!")
		os.Exit(0)
	}
	fmt.Println("启动了")
}
