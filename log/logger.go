package log

import (
	"os"

	"github.com/sirupsen/logrus"
	"github.com/sutd-statnlp/service-ladrawex/config"
	"github.com/sutd-statnlp/service-ladrawex/constant"
	"github.com/sutd-statnlp/service-ladrawex/util/fileutil"
)

var (
	logger    *logrus.Entry
	logFile   *os.File
	appConfig *config.AppConfig
)

// Init configures logger.
func init() {
	appConfig = config.Default()

	logFields := logrus.Fields{}
	logger = logrus.StandardLogger().WithFields(logFields)
	logrus.SetOutput(os.Stdout)

	if appConfig.Mode == constant.ProdMode {
		logrus.SetFormatter(&logrus.JSONFormatter{})
	}

	level, err := logrus.ParseLevel(appConfig.Log.Level)
	if err != nil {
		Error(err)
		level = logrus.DebugLevel
	}
	logrus.SetLevel(level)

	if appConfig.Log.EnableFile {
		OpenFile()
	}
}

// Info logs message at level Info.
func Info(args ...interface{}) {
	logger.Info(args...)
}

// Debug logs message at level Debug.
func Debug(args ...interface{}) {
	logger.Debug(args...)
}

// Warn logs message at level Warn.
func Warn(args ...interface{}) {
	logger.Warn(args...)
}

// Error logs message at level Error.
func Error(args ...interface{}) {
	logger.Error(args...)
}

// Fatal logs message at level Fatal and exit.
func Fatal(args ...interface{}) {
	logger.Fatal(args...)
}

// Panic logs message at level Panic and call panic.
func Panic(args ...interface{}) {
	logger.Panic(args...)
}

// OpenFile opens file for writing logging messages.
func OpenFile() bool {
	filePath := fileutil.FullPath(appConfig.Log.FilePath)
	Debug("Request to open file with path: ", filePath)
	var err error
	logFile, err = os.OpenFile(filePath, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		Error(err)
		return false
	}
	logrus.SetOutput(logFile)
	return true
}

// CloseFile closes logging file.
func CloseFile() bool {
	Debug("Request to close file: ", logFile)
	logrus.SetOutput(os.Stdout)
	err := logFile.Close()
	if err != nil {
		Error(err)
		return false
	}
	return true
}
