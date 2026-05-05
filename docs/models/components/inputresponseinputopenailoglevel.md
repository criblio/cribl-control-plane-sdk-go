# InputResponseInputOpenaiLogLevel

Collector runtime log level.

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.InputResponseInputOpenaiLogLevelError

// Open enum: custom values can be created with a direct type cast
custom := components.InputResponseInputOpenaiLogLevel("custom_value")
```


## Values

| Name                                    | Value                                   |
| --------------------------------------- | --------------------------------------- |
| `InputResponseInputOpenaiLogLevelError` | error                                   |
| `InputResponseInputOpenaiLogLevelWarn`  | warn                                    |
| `InputResponseInputOpenaiLogLevelInfo`  | info                                    |
| `InputResponseInputOpenaiLogLevelDebug` | debug                                   |
| `InputResponseInputOpenaiLogLevelSilly` | silly                                   |