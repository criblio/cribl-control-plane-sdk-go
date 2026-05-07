# SystemModeWindowsMetrics

Select the level of details for system metrics

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.SystemModeWindowsMetricsBasic

// Open enum: custom values can be created with a direct type cast
custom := components.SystemModeWindowsMetrics("custom_value")
```


## Values

| Name                               | Value                              |
| ---------------------------------- | ---------------------------------- |
| `SystemModeWindowsMetricsBasic`    | basic                              |
| `SystemModeWindowsMetricsAll`      | all                                |
| `SystemModeWindowsMetricsCustom`   | custom                             |
| `SystemModeWindowsMetricsDisabled` | disabled                           |