# CreateInputMemoryModeWindowsMetrics

Select the level of details for memory metrics

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/operations"
)

value := operations.CreateInputMemoryModeWindowsMetricsBasic

// Open enum: custom values can be created with a direct type cast
custom := operations.CreateInputMemoryModeWindowsMetrics("custom_value")
```


## Values

| Name                                          | Value                                         |
| --------------------------------------------- | --------------------------------------------- |
| `CreateInputMemoryModeWindowsMetricsBasic`    | basic                                         |
| `CreateInputMemoryModeWindowsMetricsAll`      | all                                           |
| `CreateInputMemoryModeWindowsMetricsCustom`   | custom                                        |
| `CreateInputMemoryModeWindowsMetricsDisabled` | disabled                                      |