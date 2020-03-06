/*
 * @Author: moonmist.guan
 * @Date: 2020-03-05 16:15:20
 * @LastEditTime: 2020-03-05 16:20:22
 * @FilePath: /XGBlog/app/db/es/es.go
 * @Description:
 */

package es

import (
	"XGBlog/app/config"
	"XGBlog/app/loger"
	"context"

	"github.com/olivere/elastic"
)

// Es 获取es的结构体
type Es struct {
	conn *elastic.Client
}

// Default 默认方法
func Default() *Es {
	es := &Es{}
	conn, err := elastic.NewClient(elastic.SetURL(config.Configs.EsHost))
	if err != nil {
		loger.Default().Error("es connect error:", err.Error())
		panic(err.Error())
	}
	ctx := context.Background()
	_, _, err = conn.Ping(config.Configs.EsHost).Do(ctx)
	if err != nil {
		loger.Default().Error("es ping error:", err.Error())
		panic(err.Error())
	}
	es.conn = conn
	return es
}

// GetConn 获取数据库连接
func (es *Es) GetConn() *elastic.Client {
	return es.conn
}
