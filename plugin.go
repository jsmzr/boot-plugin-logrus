package plugin

import (
	log "github.com/jsmzr/boot-log"
	plugin "github.com/jsmzr/boot-plugin"
	"github.com/spf13/viper"
)

type LogrusPlugin struct {
}

const configPrefix = "boot.logging."

var defaultConfig map[string]interface{} = map[string]interface{}{
	"enabled":      true,
	"order":        -5,
	"level":        "INFO",
	"reportCaller": false,
}

func (p *LogrusPlugin) Order() int {
	return viper.GetInt(configPrefix + "order")
}

func (p *LogrusPlugin) Load() error {
	logType := "logrus"
	log.Register(logType, &LogrusConfig{})
	return log.InitLogger(logType)
}

func (p *LogrusPlugin) Enabled() bool {
	return viper.GetBool(configPrefix + "enabled")
}

func init() {
	plugin.InitDefaultConfig(configPrefix, defaultConfig)
	plugin.Register("log", &LogrusPlugin{})
}
