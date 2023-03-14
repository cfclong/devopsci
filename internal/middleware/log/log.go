package log

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

var (
	Log = NewBeegoLogger(
		beego.AppConfig.DefaultString("log::logfile", "log/start.log"),
		beego.AppConfig.DefaultString("log::level", "1"),
		beego.AppConfig.DefaultString("log::separate", "[\"error\"]"),
	)
)

func NewBeegoLogger(logFileName, logLevel, logSeparate string) *logs.BeeLogger {
	logconfig := `{
		"filename": "` + logFileName + `",
		"level": ` + logLevel + `,
		"separate": ` + logSeparate + `
	}`

	consoleLogConfig := `{
		"level": ` + logLevel + `
	}`
	log := logs.NewLogger(1000)
	log.SetLogger(logs.AdapterMultiFile, logconfig)
	log.SetLogger(logs.AdapterConsole, consoleLogConfig)
	log.EnableFuncCallDepth(true)
	log.Async()
	return log
}
