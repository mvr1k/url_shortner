package logger

import (
	"fmt"
	"go.uber.org/zap"
)

func Logger() *zap.SugaredLogger {
	logger, err := zap.NewProduction()
	if err != nil {
		panic("Error while Creating the logger")
	}
	sugar := logger.Sugar()
	fmt.Println("logger instance created :) ")
	return sugar
}
