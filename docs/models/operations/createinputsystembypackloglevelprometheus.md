# CreateInputSystemByPackLogLevelPrometheus

Collector runtime log level

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/operations"
)

value := operations.CreateInputSystemByPackLogLevelPrometheusError

// Open enum: custom values can be created with a direct type cast
custom := operations.CreateInputSystemByPackLogLevelPrometheus("custom_value")
```


## Values

| Name                                             | Value                                            |
| ------------------------------------------------ | ------------------------------------------------ |
| `CreateInputSystemByPackLogLevelPrometheusError` | error                                            |
| `CreateInputSystemByPackLogLevelPrometheusWarn`  | warn                                             |
| `CreateInputSystemByPackLogLevelPrometheusInfo`  | info                                             |
| `CreateInputSystemByPackLogLevelPrometheusDebug` | debug                                            |