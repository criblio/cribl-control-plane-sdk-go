# CreateInputCPUModeWindowsMetrics

Select the level of details for CPU metrics

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/operations"
)

value := operations.CreateInputCPUModeWindowsMetricsBasic

// Open enum: custom values can be created with a direct type cast
custom := operations.CreateInputCPUModeWindowsMetrics("custom_value")
```


## Values

| Name                                       | Value                                      |
| ------------------------------------------ | ------------------------------------------ |
| `CreateInputCPUModeWindowsMetricsBasic`    | basic                                      |
| `CreateInputCPUModeWindowsMetricsAll`      | all                                        |
| `CreateInputCPUModeWindowsMetricsCustom`   | custom                                     |
| `CreateInputCPUModeWindowsMetricsDisabled` | disabled                                   |