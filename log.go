package plugin

import (
	"strings"

	config "github.com/jsmzr/boot-config"
	log "github.com/jsmzr/boot-log"
	"github.com/sirupsen/logrus"
)

type LogrusConfig struct{}

type LogrusContainer struct {
	logger *logrus.Logger
}

func (c *LogrusConfig) Load() (log.Logger, error) {
	level, ok := config.Get("boot.logging.level")
	logger := logrus.New()
	if ok {
		levelStr := strings.ToUpper(level.Str)
		switch levelStr {
		case "DEBUG":
			logger.SetLevel(logrus.DebugLevel)
		case "INFO":
			logger.SetLevel(logrus.InfoLevel)
		case "WARN":
			logger.SetLevel(logrus.WarnLevel)
		case "ERROR":
			logger.SetLevel(logrus.ErrorLevel)
		}
	} else {
		logger.SetLevel(logrus.InfoLevel)
	}
	var format logrus.TextFormatter
	if config.Resolve("boot.logging.format", &format) == nil {
		logger.SetFormatter(&format)
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
