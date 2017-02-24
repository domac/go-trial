package logging

import (
	"github.com/uber-go/zap"
)

var mylogger zap.Logger

// Logger returns the logger instance
func Logger() zap.Logger {
	if mylogger == nil {
		mylogger = zap.New(zap.NewJSONEncoder())
	}
	return mylogger
}
