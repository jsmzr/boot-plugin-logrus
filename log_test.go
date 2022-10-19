package plugin

import "testing"

func TestLoadLog(t *testing.T) {
	logrusConfig := LogrusConfig{}
	logger, err := logrusConfig.Load()
	if err != nil {
		t.Fatal(err)
	}
	if logger == nil {
		t.Fatal("logger should not be nil")
	}
	logger.Debug("debug")
	logger.Info("info")
	logger.Warn("warn")
	logger.Error("error")
}
