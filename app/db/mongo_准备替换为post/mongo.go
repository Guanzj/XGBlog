/*
 * @Author: moonmist.guan
 * @Date: 2020-03-05 16:15:20
 * @LastEditTime: 2020-03-06 23:09:22
 * @FilePath: /XGBlog/app/db/mongo_准备替换为post/mongo.go
 * @Description:
 */

package mongo

import (
	"XGBlog/app/config"
	"XGBlog/app/loger"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// Mongo 结构体
type Mongo struct {
	conn *mongo.Client
}

// Default 默认方法
func Default() *Mongo {
	mg := &Mongo{}
	ctx, _ := context.WithTimeout(context.Background(), time.Duration(config.Configs.MongoConnTimeout)*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(config.Configs.MongoHost))
	if err != nil {
		loger.Default().Error("mongo conn error:", err.Error())
		panic(err.Error())
	}
	ctx, _ = context.WithTimeout(context.Background(), time.Duration(config.Configs.MongoConnTimeout)*time.Second)
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		loger.Default().Error("mongo ping error:", err.Error())
		panic(err.Error())
	}
	mg.conn = client
	return mg
}

// GetConn 获取数据库连接
func (m *Mongo) GetConn() *mongo.Client {
	return m.conn
}
