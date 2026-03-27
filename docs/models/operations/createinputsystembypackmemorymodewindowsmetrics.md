# CreateInputSystemByPackMemoryModeWindowsMetrics

Select the level of details for memory metrics

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/operations"
)

value := operations.CreateInputSystemByPackMemoryModeWindowsMetricsBasic

// Open enum: custom values can be created with a direct type cast
custom := operations.CreateInputSystemByPackMemoryModeWindowsMetrics("custom_value")
```


## Values

| Name                                                      | Value                                                     |
| --------------------------------------------------------- | --------------------------------------------------------- |
| `CreateInputSystemByPackMemoryModeWindowsMetricsBasic`    | basic                                                     |
| `CreateInputSystemByPackMemoryModeWindowsMetricsAll`      | all                                                       |
| `CreateInputSystemByPackMemoryModeWindowsMetricsCustom`   | custom                                                    |
| `CreateInputSystemByPackMemoryModeWindowsMetricsDisabled` | disabled                                                  |