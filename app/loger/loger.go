package loger

// Package log 基础日志组件
import (
	"XGBlog/app/config"
	"XGBlog/app/util"
	"os"
	"time"

	"github.com/k0kubun/pp"
	"github.com/mattn/go-isatty"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

func init() {
	setLevel()
	//initPP()
}

// Loger 日志指针
var Loger *logrus.Logger

// Default 日志记录器
func Default() *logrus.Logger {
	if Loger != nil {
		return Loger
	}
	appDir := util.GetAppDir()
	today := time.Now().Format("2006-01-02")
	pathMap := lfshook.PathMap{
		logrus.InfoLevel:  appDir + "/log/info-" + today + ".log",
		logrus.ErrorLevel: appDir + "/log/error-" + today + ".log",
	}

	Loger = logrus.New()

	Loger.Hooks.Add(lfshook.NewHook(
		pathMap,
		&logrus.JSONFormatter{},
	))
	return Loger
}

func initPP() {
	out := os.Stdout
	pp.SetDefaultOutput(out)

	if !isatty.IsTerminal(out.Fd()) {
		pp.ColoringEnabled = false
	}
}

var levels = map[string]logrus.Level{
	"panic": logrus.PanicLevel,
	"fatal": logrus.FatalLevel,
	"error": logrus.ErrorLevel,
	"warn":  logrus.WarnLevel,
	"info":  logrus.InfoLevel,
	"debug": logrus.DebugLevel,
}

func setLevel() {
	levelConf := config.Configs.LogLevel

	if levelConf == "" {
		levelConf = "info"
	}

	if level, ok := levels[levelConf]; ok {
		logrus.SetLevel(level)
	} else {
		logrus.SetLevel(logrus.DebugLevel)
	}
}

// PP 类似 PHP 的 var_dump
func PP(args ...interface{}) {
	pp.Println(args...)
}
