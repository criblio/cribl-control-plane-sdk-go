# InputPrometheusLogLevel

Collector runtime log level

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.InputPrometheusLogLevelError

// Open enum: custom values can be created with a direct type cast
custom := components.InputPrometheusLogLevel("custom_value")
```


## Values

| Name                           | Value                          |
| ------------------------------ | ------------------------------ |
| `InputPrometheusLogLevelError` | error                          |
| `InputPrometheusLogLevelWarn`  | warn                           |
| `InputPrometheusLogLevelInfo`  | info                           |
| `InputPrometheusLogLevelDebug` | debug                          |