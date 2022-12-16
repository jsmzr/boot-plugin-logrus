package plugin

import (
	"fmt"
	"strings"

	log "github.com/jsmzr/boot-log"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type LogrusConfig struct{}

type LogrusContainer struct {
	logger *logrus.Logger
}

func (c *LogrusConfig) Load() (log.Logger, error) {
	level := viper.GetString(configPrefix + ".level")
	logger := logrus.New()
	switch strings.ToUpper(level) {
	case "DEBUG":
		logger.SetLevel(logrus.DebugLevel)
	case "WARN":
		logger.SetLevel(logrus.WarnLevel)
	case "ERROR":
		logger.SetLevel(logrus.ErrorLevel)
	default:
		logger.SetLevel(logrus.InfoLevel)
	}
	var format logrus.TextFormatter
	if err := viper.UnmarshalKey(configPrefix+".format", &format); err == nil {
		logger.SetFormatter(&format)
	} else {
		fmt.Printf("unmarshalKey [boot.loggin.format] failed, %v \n", err)
	}
	return &LogrusContainer{
		logger: logger,
	}, nil
}

func (c *LogrusContainer) Debug(msg string, params ...interface{}) {
	c.logger.Debugf(msg, params...)
}
func (c *LogrusContainer) Info(msg string, params ...interface{}) {
	c.logger.Infof(msg, params...)
}
func (c *LogrusContainer) Warn(msg string, params ...interface{}) {
	c.logger.Warnf(msg, params...)
}
func (c *LogrusContainer) Error(msg string, params ...interface{}) {
	c.logger.Errorf(msg, params...)
}
