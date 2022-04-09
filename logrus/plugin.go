package log

import (
	_ "github.com/jsmzr/bootstrap-log-logrus/logrus"
	"github.com/jsmzr/bootstrap-log/log"
	"github.com/jsmzr/bootstrap-plugin/plugin"
)

type LogrusPlugin struct {
}

func (p *LogrusPlugin) Load() error {
	// TODO 日志配置
	return log.InitLogger("logrus")
}

func init() {
	plugin.Register("log", &LogrusPlugin{})
}
