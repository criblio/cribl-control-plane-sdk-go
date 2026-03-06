# CreateInputDiskModeWindowsMetrics

Select the level of details for disk metrics

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/operations"
)

value := operations.CreateInputDiskModeWindowsMetricsBasic

// Open enum: custom values can be created with a direct type cast
custom := operations.CreateInputDiskModeWindowsMetrics("custom_value")
```


## Values

| Name                                        | Value                                       |
| ------------------------------------------- | ------------------------------------------- |
| `CreateInputDiskModeWindowsMetricsBasic`    | basic                                       |
| `CreateInputDiskModeWindowsMetricsAll`      | all                                         |
| `CreateInputDiskModeWindowsMetricsCustom`   | custom                                      |
| `CreateInputDiskModeWindowsMetricsDisabled` | disabled                                    |