package initialize

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	hertzlogrus "github.com/hertz-contrib/obs-opentelemetry/logging/logrus"
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
	"msp/gateway/config"
	"os"
	"path"
	"time"
)

// InitLogger to init logrus
func InitLogger() {
	// Customizable output directory.
	logConfig := config.GlobalLocalConfig.LogConfig
	logFilePath := path.Join(logConfig.LogDir, config.GlobalLocalConfig.Name)
	if err := os.MkdirAll(logFilePath, 0o777); err != nil {
		panic(err)
	}

	// Set filename to date
	logFileName := time.Now().Format("2006-01-02") + ".log"
	fileName := path.Join(logFilePath, logFileName)
	if _, err := os.Stat(fileName); err != nil {
		if _, err := os.Create(fileName); err != nil {
			panic(err)
		}
	}

	logger := hertzlogrus.WithLogger(
		logrus.StandardLogger(),
	)

	// Provides compression and deletion
	lumberjackLogger := &lumberjack.Logger{
		Filename:   fileName,
		MaxSize:    20,   // A file can be up to 20M.
		MaxBackups: 5,    // Save up to 5 files at the same time.
		MaxAge:     10,   // A file can exist for a maximum of 10 days.
		Compress:   true, // Compress with gzip.
	}
	if logConfig.LogDebug || config.IsDev {
		logger.SetLevel(hlog.LevelDebug)
	} else {
		logger.SetLevel(hlog.LevelWarn)
	}
	logger.SetOutput(lumberjackLogger)

	hlog.SetLogger(logger)
}
