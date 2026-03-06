# CreateInputSystemByPackDiskModeWindowsMetrics

Select the level of details for disk metrics

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/operations"
)

value := operations.CreateInputSystemByPackDiskModeWindowsMetricsBasic

// Open enum: custom values can be created with a direct type cast
custom := operations.CreateInputSystemByPackDiskModeWindowsMetrics("custom_value")
```


## Values

| Name                                                    | Value                                                   |
| ------------------------------------------------------- | ------------------------------------------------------- |
| `CreateInputSystemByPackDiskModeWindowsMetricsBasic`    | basic                                                   |
| `CreateInputSystemByPackDiskModeWindowsMetricsAll`      | all                                                     |
| `CreateInputSystemByPackDiskModeWindowsMetricsCustom`   | custom                                                  |
| `CreateInputSystemByPackDiskModeWindowsMetricsDisabled` | disabled                                                |