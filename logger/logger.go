package logger

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"sync"
)

var (
	Once         sync.Once
	globalLogger *logrus.Logger
)

type Logger struct {
	LogLevel string `toml:"log_level"`
}

func LogInstance() *logrus.Logger {
	Once.Do(func() {
		globalLogger = logrus.New()
		fmt.Println("Create Global logger")
	})

	return globalLogger
}
