# CreateInputSystemModeWindowsMetrics

Select the level of details for system metrics

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/operations"
)

value := operations.CreateInputSystemModeWindowsMetricsBasic

// Open enum: custom values can be created with a direct type cast
custom := operations.CreateInputSystemModeWindowsMetrics("custom_value")
```


## Values

| Name                                          | Value                                         |
| --------------------------------------------- | --------------------------------------------- |
| `CreateInputSystemModeWindowsMetricsBasic`    | basic                                         |
| `CreateInputSystemModeWindowsMetricsAll`      | all                                           |
| `CreateInputSystemModeWindowsMetricsCustom`   | custom                                        |
| `CreateInputSystemModeWindowsMetricsDisabled` | disabled                                      |