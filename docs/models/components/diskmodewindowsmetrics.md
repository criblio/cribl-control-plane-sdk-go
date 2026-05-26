# DiskModeWindowsMetrics

Select the level of details for disk metrics

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.DiskModeWindowsMetricsBasic

// Open enum: custom values can be created with a direct type cast
custom := components.DiskModeWindowsMetrics("custom_value")
```


## Values

| Name                             | Value                            |
| -------------------------------- | -------------------------------- |
| `DiskModeWindowsMetricsBasic`    | basic                            |
| `DiskModeWindowsMetricsAll`      | all                              |
| `DiskModeWindowsMetricsCustom`   | custom                           |
| `DiskModeWindowsMetricsDisabled` | disabled                         |