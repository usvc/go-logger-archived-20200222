/*
Package logger is a standardised logger for services in the `usvc` namespace.

Global usage:

	package x

	import "github.com/usvc/go-logger"

	var log = logger.New()

	func main() {
		log.Info("hello world")
	}

Module usage:

	package x

	import "github.com/usvc/go-logger"

	func NewX() *Something {
		config := logger.NewConfig()
		config.Fields["module"] = "NewX"
		something := &Something{
			logger: logger.New(config)
		}
		logger.Info("initialised something")
		return something
	}

For documentation, check out: https://github.com/usvc/go-logger
*/
package logger
