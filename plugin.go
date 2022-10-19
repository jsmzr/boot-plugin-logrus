package plugin

import (
	log "github.com/jsmzr/bootstrap-log"
	plugin "github.com/jsmzr/bootstrap-plugin"
)

type LogrusPlugin struct {
}

func (p *LogrusPlugin) Order() int {
	return 0
}

func (p *LogrusPlugin) Load() error {
	logType := "logrus"
	log.Register(logType, &LogrusConfig{})
	return log.InitLogger(logType)
}

func init() {
	plugin.Register("log", &LogrusPlugin{})
}
