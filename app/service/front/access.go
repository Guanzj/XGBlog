/*
 * @Author: moonmist.guan
 * @Date: 2020-03-05 15:08:59
 * @LastEditTime: 2020-03-05 15:24:22
 * @FilePath: /XGBlog/app/service/front/access.go
 * @Description:
 */

package front

import (
	"errors"
	"fmt"
)

// AccessLog 通过日志的结构体
type AccessLog struct {
}

// AccessLogParams 通过日志的参数
type AccessLogParams struct {
	IP        string `json:"ip"`
	Timestamp int64  `json:"timestamp"`
	Path      string `json:"path"`
	Date      string `json:"date"`
}

// Add 通过日志结构体的方法，接收日志参数列表作为入参，写入MongoDB数据库，需要注意
func (a *AccessLog) Add(log *AccessLogParams) error {
	fmt.Println("准备写往数据库")
	fmt.Println(log)
	// mongoconn := mongo.Default().GetConn()
	// collection := mongoconn.Database("myblog").Collection("access_log")
	// ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	// accessLog, err := bson.Marshal(log)
	// if err != nil {
	// 	return err
	// }
	// _, err = collection.InsertOne(ctx, accessLog)
	// return err
	return errors.New("先凑活用的 接入日志记录的错误返回")
}
