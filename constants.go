package logger

import "github.com/sirupsen/logrus"

// Level defines the symbol for log levels
type Level string

// Format defines the symbol for log format
type Format string

const (
	// FormatJSON defines the symbol for using the JSON format
	FormatJSON Format = "FORMAT_JSON"
	// FormatText defines the symbol for using the text format
	FormatText Format = "FORMAT_TEXT"
	// LevelTrace defines the symbol for random stuff
	LevelTrace Level = "LEVEL_TRACE"
	// LevelDebug defines the symbol for good to know stuff
	LevelDebug Level = "LEVEL_DEBUG"
	// LevelInfo defines the symbol for need to know stuff
	LevelInfo Level = "LEVEL_INFO"
	// LevelWarn defines the symbol for when something might've fucked up
	LevelWarn Level = "LEVEL_WARN"
	// LevelError defines the symbol for when shit hits the fan
	LevelError Level = "LEVEL_ERROR"
	// TimestampFormat defines the timestamp format for the logger
	TimestampFormat string = "2006-01-02 15:04:05"
	// DataKey defines the key name for extra data when using the JSON format
	DataKey string = "@data"
)

// FieldMap defines the field mapping for the logger's default fields
var FieldMap = logrus.FieldMap{
	logrus.FieldKeyFile:  "@file",
	logrus.FieldKeyFunc:  "@caller",
	logrus.FieldKeyLevel: "@level",
	logrus.FieldKeyMsg:   "@message",
	logrus.FieldKeyTime:  "@timestamp",
}

// Levels provides a mapping to log levels so that we can
// substitute out the values easily if we switch the base logger
var Levels = map[Level]logrus.Level{
	LevelTrace: logrus.TraceLevel,
	LevelDebug: logrus.DebugLevel,
	LevelInfo:  logrus.InfoLevel,
	LevelWarn:  logrus.WarnLevel,
	LevelError: logrus.ErrorLevel,
}
