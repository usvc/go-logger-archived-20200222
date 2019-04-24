# `usvc/go-logger`
A sensible standardised logger for services in `usvc`.

# Usage

## Import

```golang
import "github.com/usvc/go-logger"
```

## Usage

### Global Logger

```golang
package main

import "github.com/usvc/go-logger"

var log = logger.New()

func main() {
  log.Info("hello world")
}
```

### Module Logger

```golang
package main

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

type Something struct{}
```

## Configuration

### Via Code

Available fields are:

| Property | Type | Description |
| --- | --- | --- |
| `Fields` | `map[string]interface{}` | Defines the fields logged with the logs |
| `Format` | `string` | Defines the format - JSON or text |
| `Level` | `string` | Defines the level - trace, debug, info, warn, or error |
| `OutputTo` | `io.Writer` | Defines the output, set this to pipe to a `byte.Buffer` for example |
| `Debug` | `bool` | Defines whether to run in debug mode |

### Via Environment

| Key | Description | Values |
| --- | --- | --- |
| `FORMAT` | Defines the log format | `FORMAT_JSON`, `FORMAT_TEXT` |
| `LEVEL` | Defines the log level | `LEVEL_TRACE`, `LEVEL_DEBUG`, `LEVEL_INFO`, `LEVEL_WARN`, `LEVEL_ERROR` |
| `DEBUG` | Defines whether to run in debug mode | `1`, `0` |

# License
This project is licensed under the MIT license. See the [LICENSE](./LICENSE) file for the full text.

# Cheers
