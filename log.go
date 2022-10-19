package plugin

import (
	"fmt"
	"strings"

	config "github.com/jsmzr/bootstrap-config"
	log "github.com/jsmzr/bootstrap-log"
	"github.com/sirupsen/logrus"
)

type LogrusConfig struct{}

type LogrusContainer struct {
	logger *logrus.Logger
}

func (c *LogrusConfig) Load() (log.Logger, error) {
	level, ok := config.Get("logging.level")
	logger := logrus.New()
	if ok {
		levelStr := strings.ToUpper(level.Str)
		fmt.Printf("[Bootstrap-Plugin-Logrus]  Set logger level [%v]\n", levelStr)
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
	err := config.Resolve("logging.format", &format)
	if err != nil {
		logger.Printf("[Bootstrap-Plugin-Logrus]  Get logger format error: %s \n", err.Error())
		logger.Println("[Bootstrap-Plugin-Logrus]  Use default format")
	} else {
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
