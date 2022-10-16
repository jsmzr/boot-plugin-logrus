package plugin

import (
	"github.com/jsmzr/bootstrap-log/log"
	"github.com/jsmzr/bootstrap-plugin/plugin"
)

type LogrusPlugin struct {
}

func (p *LogrusPlugin) Order() int {
	return 0
}

func (p *LogrusPlugin) Load() error {
	// TODO 日志配置
	logType := "logrus"
	log.Register(logType, &LogrusConfig{})
	return log.InitLogger(logType)
}

func init() {
	plugin.Register("log", &LogrusPlugin{})
}
