package logger

import (
	"io"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

const (
	// DefaultLevel defines the default level of logging
	DefaultLevel Level = LevelTrace
	// DefaultFormat defines the default log formatting
	DefaultFormat Format = FormatText
	// DefaultDebug defines the default debug value
	DefaultDebug = false
	// EnvKeyFormat defines the environment key for the formatter
	EnvKeyFormat = "FORMAT"
	// EnvKeyLevel defines the environment key for the log level
	EnvKeyLevel = "LEVEL"
	// EnvKeyDebug defines the environment key for the debug definition
	EnvKeyDebug = "DEBUG"
)

// NewConfig is a convenience function for creating a logger
// configuration with default values
func NewConfig() *Config {
	return &Config{
		Fields:   make(map[string]interface{}),
		Format:   DefaultFormat,
		Level:    DefaultLevel,
		OutputTo: nil,
		Debug:    DefaultDebug,
	}
}

// NewConfigFromEnvironment is a convenience function for retrieving
// the logger configuration from the environment. Use the :prefix to
// define an environment variable prefix:
//
// Default config (will draw values from FORMAT for example)
//   config := NewConfigFromEnvironment()
//
// Environment config (will draw values from LOGGER_FORMAT for example)
//   config := NewConfigFromEnvironment("LOGGER")
func NewConfigFromEnvironment(prefix ...string) *Config {
	if len(prefix) > 0 {
		viper.SetEnvPrefix(prefix[0])
	}
	return &Config{
		Fields:   make(map[string]interface{}),
		Format:   Format(stringFromEnv(EnvKeyFormat, string(DefaultFormat))),
		Level:    Level(stringFromEnv(EnvKeyLevel, string(DefaultLevel))),
		OutputTo: nil,
		Debug:    boolFromEnv(EnvKeyDebug, DefaultDebug),
	}
}

func boolFromEnv(key string, defaultValue bool) bool {
	viper.SetDefault(key, defaultValue)
	viper.BindEnv(key)
	return viper.GetBool(key)
}
func stringFromEnv(key string, defaultValue string) string {
	viper.SetDefault(key, defaultValue)
	viper.BindEnv(key)
	return viper.GetString(key)
}

// Config holds the... (surprise, surprise) configuration!
type Config struct {
	Fields   map[string]interface{}
	Format   Format
	Level    Level
	OutputTo io.Writer
	Debug    bool
}

// GetFormatter returns a
func (config *Config) GetFormatter() logrus.Formatter {
	if config.Format == FormatJSON {
		return &logrus.JSONFormatter{
			DataKey:         DataKey,
			FieldMap:        FieldMap,
			PrettyPrint:     config.Debug,
			TimestampFormat: TimestampFormat,
		}
	}
	return &logrus.TextFormatter{
		DisableSorting:   true,
		FieldMap:         FieldMap,
		ForceColors:      true,
		FullTimestamp:    true,
		QuoteEmptyFields: true,
		TimestampFormat:  TimestampFormat,
	}
}

// GetLevel retrieves the minimum level at which we'll
// be displaying logs
func (config *Config) GetLevel() logrus.Level {
	return Levels[config.Level]
}
