# NetworkModeWindowsMetrics

Select the level of details for network metrics

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.NetworkModeWindowsMetricsBasic

// Open enum: custom values can be created with a direct type cast
custom := components.NetworkModeWindowsMetrics("custom_value")
```


## Values

| Name                                | Value                               |
| ----------------------------------- | ----------------------------------- |
| `NetworkModeWindowsMetricsBasic`    | basic                               |
| `NetworkModeWindowsMetricsAll`      | all                                 |
| `NetworkModeWindowsMetricsCustom`   | custom                              |
| `NetworkModeWindowsMetricsDisabled` | disabled                            |