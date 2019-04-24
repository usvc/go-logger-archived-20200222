package logger

import "github.com/sirupsen/logrus"

// New creates and returns a new logger
func New(config *Config) Logger {
	logger := &logrus.Logger{
		Formatter:    config.GetFormatter(),
		Level:        config.GetLevel(),
		Out:          config.OutputTo,
		ReportCaller: config.Debug,
	}
	log := logger.WithFields(config.Fields)
	return log
}

// Logger interface to fulfil
type Logger interface {
	Tracef(format string, data ...interface{})
	Trace(data ...interface{})
	Debugf(format string, data ...interface{})
	Debug(data ...interface{})
	Infof(format string, data ...interface{})
	Info(data ...interface{})
	Warnf(format string, data ...interface{})
	Warn(data ...interface{})
	Errorf(format string, data ...interface{})
	Error(data ...interface{})
	Fatalf(format string, data ...interface{})
	Fatal(data ...interface{})
	Panicf(format string, data ...interface{})
	Panic(data ...interface{})
	WithFields(logrus.Fields) *logrus.Entry
}
