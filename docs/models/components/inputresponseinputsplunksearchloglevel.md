# InputResponseInputSplunkSearchLogLevel

Collector runtime log level (verbosity)

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.InputResponseInputSplunkSearchLogLevelError

// Open enum: custom values can be created with a direct type cast
custom := components.InputResponseInputSplunkSearchLogLevel("custom_value")
```


## Values

| Name                                          | Value                                         |
| --------------------------------------------- | --------------------------------------------- |
| `InputResponseInputSplunkSearchLogLevelError` | error                                         |
| `InputResponseInputSplunkSearchLogLevelWarn`  | warn                                          |
| `InputResponseInputSplunkSearchLogLevelInfo`  | info                                          |
| `InputResponseInputSplunkSearchLogLevelDebug` | debug                                         |