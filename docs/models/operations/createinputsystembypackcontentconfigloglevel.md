# CreateInputSystemByPackContentConfigLogLevel

Collector runtime log level.

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/operations"
)

value := operations.CreateInputSystemByPackContentConfigLogLevelError

// Open enum: custom values can be created with a direct type cast
custom := operations.CreateInputSystemByPackContentConfigLogLevel("custom_value")
```


## Values

| Name                                                | Value                                               |
| --------------------------------------------------- | --------------------------------------------------- |
| `CreateInputSystemByPackContentConfigLogLevelError` | error                                               |
| `CreateInputSystemByPackContentConfigLogLevelWarn`  | warn                                                |
| `CreateInputSystemByPackContentConfigLogLevelInfo`  | info                                                |
| `CreateInputSystemByPackContentConfigLogLevelDebug` | debug                                               |
| `CreateInputSystemByPackContentConfigLogLevelSilly` | silly                                               |