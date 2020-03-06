/*
 * @Author: moonmist.guan
 * @Date: 2020-03-05 16:15:20
 * @LastEditTime: 2020-03-06 00:00:07
 * @FilePath: /XGBlog/app/db/mysql/mysql.go
 * @Description:
 */

package mysql

import (
	"XGBlog/app/config"
	"XGBlog/app/loger"
	"fmt"

	"github.com/jinzhu/gorm"

	//这一行需要保留，否则会报import _ "github.com/go-sql-driver/mysql"错误
	_ "github.com/go-sql-driver/mysql"
)

// Mysql gorm结构体
type Mysql struct {
	conn *gorm.DB
}

// Default 默认方法
func Default() (db *Mysql) {
	mysql := &Mysql{}
	conn, err := gorm.Open(config.Configs.DBDriver, config.Configs.DBUser+":"+config.Configs.DBPass+"@tcp("+config.Configs.DBHost+":"+config.Configs.DBPort+")/"+config.Configs.DBName)
	if err != nil {
		loger.Default().Error("mysql connect error:", err.Error()+config.Configs.DBDriver)
		panic(err.Error() + config.Configs.DBDriver)
	}
	if config.Configs.DBDebug {
		conn = conn.Debug()
	}
	mysql.conn = conn
	return mysql
}

// GetConn 获取数据库连接
func (m *Mysql) GetConn() *gorm.DB {
	fmt.Println("获取数据库连接成功")
	return m.conn
}
