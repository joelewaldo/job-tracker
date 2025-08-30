package logger

import (
	"os"
	"strings"

	"github.com/joelewaldo/job-tracker/api/internal/config"
	"github.com/sirupsen/logrus"
)

var Log = logrus.New()

func Init(cfg *config.Config) {
	Log.Out = os.Stdout
	Log.SetFormatter(&logrus.JSONFormatter{})

	level, err := logrus.ParseLevel(strings.ToLower(cfg.LogLevel))
	if err != nil {
		Log.Warnf("invalid LOG_LEVEL '%s', defaulting to info", cfg.LogLevel)
		level = logrus.InfoLevel
	}
	Log.SetLevel(level)
}

