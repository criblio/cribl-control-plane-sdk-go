# TelemetryType

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.TelemetryTypeLogs

// Open enum: custom values can be created with a direct type cast
custom := components.TelemetryType("custom_value")
```


## Values

| Name                   | Value                  |
| ---------------------- | ---------------------- |
| `TelemetryTypeLogs`    | logs                   |
| `TelemetryTypeMetrics` | metrics                |