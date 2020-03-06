/*
 * @Author: moonmist.guan
 * @Date: 2020-03-05 14:33:51
 * @LastEditTime: 2020-03-06 23:38:02
 * @FilePath: /XGBlog/app/router/router.go
 * @Description:
 */

package router

import (
	"XGBlog/app/controller/admin"
	"XGBlog/app/controller/front"
	"XGBlog/app/middleware"
	"XGBlog/app/protocol"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Router 通过gin生成路由结构体
type Router struct {
	engine *gin.Engine
}

// Default 通过gin获得路由的default
func Default() *Router {
	router := &Router{}
	router.engine = gin.Default()
	return router
}

// Run 方法是Router结构体的方法
func (r *Router) Run() {
	r.SetAccessLog()
	r.SetCors()
	r.setFront()
	r.setAdmin()
	r.set404()
}

// GetEngin 返回结构体的engine
func (r *Router) GetEngin() *gin.Engine {
	return r.engine
}

// SetAccessLog 调用中间件的日志
func (r *Router) SetAccessLog() {
	r.engine.Use(middleware.AddAccessLog())
}

// SetCors 设置跨域
func (r *Router) SetCors() {
	r.engine.Use(middleware.Cors())
}

// setFront 设置前台博客页面的路由--外部不可见
// 该处调用的是controller中的方法
func (r *Router) setFront() {
	articleFrontCtrl := front.Articles{}
	r.engine.GET("/articles", articleFrontCtrl.GetList)
	r.engine.GET("/categories", articleFrontCtrl.GetCategories)
	r.engine.GET("/article/:id", articleFrontCtrl.Show)

}

// setAdmin 设置后台管理页面的路由--外部不可见
func (r *Router) setAdmin() {
	//后台管理
	loginAdminCtrl := admin.Login{}
	articleAdminCtrl := admin.Articles{}
	userAdminCtrl := admin.User{}
	r.engine.POST("/adapi/login", loginAdminCtrl.Login)
	authorized := r.engine.Group("/adapi")
	authorized.Use(middleware.CheckAuth())
	{
		authorized.POST("/article", articleAdminCtrl.Add)          //ok
		authorized.PUT("/article/:id", articleAdminCtrl.Update)    //ok
		authorized.GET("/articles", articleAdminCtrl.GetList)      //ok
		authorized.DELETE("/article/:id", articleAdminCtrl.Delete) //ok
		authorized.GET("/article/:id", articleAdminCtrl.Show)      //ok
		authorized.GET("/user", userAdminCtrl.Show)                //ok
	}
}

// set404 设置404页面的路由--外部不可见
func (r *Router) set404() {
	r.engine.NoRoute(func(context *gin.Context) {
		resp := protocol.Resp{Ret: 404, Msg: "page not exists!", Data: ""}
		//返回404状态码
		context.JSON(http.StatusNotFound, resp)
	})
}
