# InputSplunkSearchLogLevel

Collector runtime log level (verbosity)

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.InputSplunkSearchLogLevelError

// Open enum: custom values can be created with a direct type cast
custom := components.InputSplunkSearchLogLevel("custom_value")
```


## Values

| Name                             | Value                            |
| -------------------------------- | -------------------------------- |
| `InputSplunkSearchLogLevelError` | error                            |
| `InputSplunkSearchLogLevelWarn`  | warn                             |
| `InputSplunkSearchLogLevelInfo`  | info                             |
| `InputSplunkSearchLogLevelDebug` | debug                            |