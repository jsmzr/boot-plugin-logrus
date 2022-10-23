package plugin

import (
	log "github.com/jsmzr/boot-log"
	plugin "github.com/jsmzr/boot-plugin"
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

func (p *LogrusPlugin) Enabled() bool {
	return true
}

func init() {
	plugin.Register("log", &LogrusPlugin{})
}
