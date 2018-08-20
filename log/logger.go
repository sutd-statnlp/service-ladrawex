package log

import (
	"os"

	"github.com/sirupsen/logrus"
	"github.com/sutd-statnlp/service-ladrawex/config"
)

var (
	logger    *logrus.Entry
	logFile   *os.File
	logConfig *config.LogConfig
)

// Init configures logger.
func init() {
	logConfig := config.Default().Log

	logFields := logrus.Fields{}
	logger = logrus.StandardLogger().WithFields(logFields)
	logrus.SetOutput(os.Stdout)

	level, err := logrus.ParseLevel(logConfig.Level)
	if err != nil {
		Error(err)
		level = logrus.DebugLevel
	}
	logrus.SetLevel(level)

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
	filePath := logConfig.FilePath
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
	if logFile == nil {
		return false
	}
	logrus.SetOutput(os.Stdout)
	err := logFile.Close()
	if err != nil {
		Error(err)
		return false
	}
	return true
}
