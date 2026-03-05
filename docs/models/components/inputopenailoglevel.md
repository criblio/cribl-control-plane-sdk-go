# InputOpenaiLogLevel

Collector runtime log level.

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.InputOpenaiLogLevelError

// Open enum: custom values can be created with a direct type cast
custom := components.InputOpenaiLogLevel("custom_value")
```


## Values

| Name                       | Value                      |
| -------------------------- | -------------------------- |
| `InputOpenaiLogLevelError` | error                      |
| `InputOpenaiLogLevelWarn`  | warn                       |
| `InputOpenaiLogLevelInfo`  | info                       |
| `InputOpenaiLogLevelDebug` | debug                      |
| `InputOpenaiLogLevelSilly` | silly                      |