package plugin

import (
	"os"
	"testing"

	"github.com/jsmzr/boot-plugin-logrus/config"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func TestLoadLevelLog(t *testing.T) {
	logrusConfig := LogrusConfig{}
	// default level
	if _, err := logrusConfig.Load(); err != nil {
		t.Fatal(err)
	}

	// debug level
	viper.Set("boot.logging.level", "DEBUG")
	if _, err := logrusConfig.Load(); err != nil {
		t.Fatal(err)
	}

	// warn level
	viper.Set("boot.logging.level", "WARN")
	if _, err := logrusConfig.Load(); err != nil {
		t.Fatal(err)
	}

	// error level
	viper.Set("boot.logging.level", "ERROR")
	if _, err := logrusConfig.Load(); err != nil {
		t.Fatal(err)
	}
}

func TestLoadFormatLog(t *testing.T) {
	logrusConfig := LogrusConfig{}
	// default level
	viper.Set("boot.logging.format", "failed")
	if _, err := logrusConfig.Load(); err != nil {
		t.Fatal(err)
	}
	viper.Set("boot.logging.format.timestampFormat", "test")
	if _, err := logrusConfig.Load(); err != nil {
		t.Fatalf("load failed, %v", err)
	}

}

func TestLog(t *testing.T) {
	logrusConfig := LogrusConfig{}
	// default level
	log, err := logrusConfig.Load()
	if err != nil {
		t.Fatal(err)
	}
	if log == nil {
		t.Fatal("log instance is null")
	}
	log.Debug("debug")
	log.Info("info")
	log.Warn("war")
	log.Error("error")
}

func TestConfigLoad(t *testing.T) {
	logrusConfig := LogrusConfig{}
	config.Out = os.Stderr
	if _, err := logrusConfig.Load(); err != nil {
		t.Fatal(err)
	}
	config.Out = nil

	config.Hooks = make(logrus.LevelHooks)
	if _, err := logrusConfig.Load(); err != nil {
		t.Fatal(err)
	}
	config.Hooks = nil
	config.ExitFunc = os.Exit
	if _, err := logrusConfig.Load(); err != nil {
		t.Fatal(err)
	}

}
