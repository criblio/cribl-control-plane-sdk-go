# CreateInputSystemByPackCPUModeWindowsMetrics

Select the level of details for CPU metrics

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/operations"
)

value := operations.CreateInputSystemByPackCPUModeWindowsMetricsBasic

// Open enum: custom values can be created with a direct type cast
custom := operations.CreateInputSystemByPackCPUModeWindowsMetrics("custom_value")
```


## Values

| Name                                                   | Value                                                  |
| ------------------------------------------------------ | ------------------------------------------------------ |
| `CreateInputSystemByPackCPUModeWindowsMetricsBasic`    | basic                                                  |
| `CreateInputSystemByPackCPUModeWindowsMetricsAll`      | all                                                    |
| `CreateInputSystemByPackCPUModeWindowsMetricsCustom`   | custom                                                 |
| `CreateInputSystemByPackCPUModeWindowsMetricsDisabled` | disabled                                               |