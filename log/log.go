package log

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"path"
	"strconv"
)

var logger *logrus.Logger
var file *os.File

func Init(logFile string) {
	var err error
	file, err = createLogFile(logFile)
	if err != nil {
		fmt.Println("fail to create log file, ", err)
		os.Exit(1)
	}
	logger = logrus.New()
	logger.Formatter = &Formatter{}
	logger.Out = file
}

func createLogFile(logFile string) (*os.File, error) {
	perm, err := strconv.ParseInt("0666", 8, 64)
	fileName := logFile
	filePath := path.Dir(fileName)

	err = os.MkdirAll(filePath, os.FileMode(perm))
	if err != nil {
		return nil, err
	}

	fd, err := os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND|os.O_CREATE, os.FileMode(perm))
	if err == nil {
		_ = os.Chmod(fileName, os.FileMode(perm))
	}
	return fd, err
}

func Debug(args ...interface{}) {
	logger.Debug(args)
}

func Info(args ...interface{}) {
	logger.Info(args)
}

func Warn(args ...interface{}) {
	logger.Warn(args)
}

func Error(args ...interface{}) {
	logger.Error(args)
}

func Writer() io.Writer {
	return logger.Writer()
}
