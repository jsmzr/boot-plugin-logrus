package config

import (
	"io"

	"github.com/sirupsen/logrus"
)

var Out io.Writer
var Hooks logrus.LevelHooks
var ExitFunc func(int)
