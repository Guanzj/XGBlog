/*
 * @Author: moonmist.guan
 * @Date: 2020-03-05 15:47:59
 * @LastEditTime: 2020-03-05 15:49:23
 * @FilePath: /XGBlog/app/config/config.go
 * @Description:
 */

package config

import "gopkg.in/ini.v1"

type SysConfig struct {
	Env              string `ini:"env"`
	DBDriver         string `ini:"db_driver"`
	DBHost           string `ini:"db_host"`
	DBPort           string `ini:"db_port"`
	DBUser           string `ini:"db_user"`
	DBPass           string `ini:"db_pass"`
	DBName           string `ini:"db_name"`
	DBDebug          bool   `ini:"db_debug"`
	HttpListenPort   string `ini:"http_listen_port"`
	JwtSecret        string `ini:"jwt_secret"`
	JwtExprTime      int64  `ini:"jwt_expr_time"`
	LogDir           string `ini:"log_dir"`
	LogLevel         string `ini:"log_level"`
	GaSecret         string `ini:"ga_secret"`
	EsHost           string `ini:"es_host"`
	MongoHost        string `ini:"mongo_host"`
	MongoConnTimeout int64  `ini:"mongo_conn_timeout"`
}

// Configs 配置
var Configs *SysConfig = &SysConfig{}

//Default 加载系统配置文件
func Default(configFileName string) {
	config := &SysConfig{}
	conf, err := ini.Load(configFileName) //加载配置文件
	if err != nil {
		panic(err)
	}
	conf.BlockMode = false
	err = conf.MapTo(&config) //解析成结构体
	if err != nil {
		panic(err)
	}
	Configs = config
}
