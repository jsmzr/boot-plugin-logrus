package plugin

import (
	"fmt"
	"strings"

	"github.com/jsmzr/bootstrap-config/config"
	"github.com/jsmzr/bootstrap-log/log"
	"github.com/sirupsen/logrus"
)

type LogrusConfig struct{}

type LogrusContainer struct {
	logger *logrus.Logger
}

func (c *LogrusConfig) Load() (log.Logger, error) {
	// TODO 一些配置支持
	level, ok := config.Get("logging.level")
	logger := logrus.New()
	if ok {
		levelStr := strings.ToUpper(level.Str)
		fmt.Printf("设置日志级别为 [%v]\n", levelStr)
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
		logger.Errorf("获取不到 logrus.format 配置使用默认配置")
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
