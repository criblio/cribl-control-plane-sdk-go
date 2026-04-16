# CreateInputContentConfigLogLevel

Collector runtime log level.

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/operations"
)

value := operations.CreateInputContentConfigLogLevelError

// Open enum: custom values can be created with a direct type cast
custom := operations.CreateInputContentConfigLogLevel("custom_value")
```


## Values

| Name                                    | Value                                   |
| --------------------------------------- | --------------------------------------- |
| `CreateInputContentConfigLogLevelError` | error                                   |
| `CreateInputContentConfigLogLevelWarn`  | warn                                    |
| `CreateInputContentConfigLogLevelInfo`  | info                                    |
| `CreateInputContentConfigLogLevelDebug` | debug                                   |
| `CreateInputContentConfigLogLevelSilly` | silly                                   |