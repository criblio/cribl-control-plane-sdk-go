# CreateInputLogLevelPrometheus

Collector runtime log level

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/operations"
)

value := operations.CreateInputLogLevelPrometheusError

// Open enum: custom values can be created with a direct type cast
custom := operations.CreateInputLogLevelPrometheus("custom_value")
```


## Values

| Name                                 | Value                                |
| ------------------------------------ | ------------------------------------ |
| `CreateInputLogLevelPrometheusError` | error                                |
| `CreateInputLogLevelPrometheusWarn`  | warn                                 |
| `CreateInputLogLevelPrometheusInfo`  | info                                 |
| `CreateInputLogLevelPrometheusDebug` | debug                                |