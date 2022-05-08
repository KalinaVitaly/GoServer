package logger

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"sync"
)

var (
	once         sync.Once
	loggerLevel  string
	globalLogger *logrus.Logger
)

func SetLoggerLevel(level string) {
	loggerLevel = level
}

func LogInstance() *logrus.Logger {
	once.Do(func() {
		globalLogger = logrus.New()
		fmt.Println("Create Global logger")
	})

	return globalLogger
}
