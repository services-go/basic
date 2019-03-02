package main

import (
	"github.com/services-go/basic/mlog"
	"time"
)

func main() {
	// linux环境下，从文件读取配置
	//mlog.Logger.SetLogConf("./mlog.xml")

	mlog.SetLogConsole("trace")
	mlog.Logger.Trace("hello, services-go!")
	mlog.Logger.Debug("hello, services-go!")
	mlog.Logger.Info("hello, services-go!")
	mlog.Logger.Warn("hello, services-go!")
	mlog.Logger.Error("hello, services-go!")

	time.Sleep(time.Millisecond)
}
