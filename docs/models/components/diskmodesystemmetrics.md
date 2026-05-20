# DiskModeSystemMetrics

Select the level of detail for disk metrics

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.DiskModeSystemMetricsBasic

// Open enum: custom values can be created with a direct type cast
custom := components.DiskModeSystemMetrics("custom_value")
```


## Values

| Name                            | Value                           |
| ------------------------------- | ------------------------------- |
| `DiskModeSystemMetricsBasic`    | basic                           |
| `DiskModeSystemMetricsAll`      | all                             |
| `DiskModeSystemMetricsCustom`   | custom                          |
| `DiskModeSystemMetricsDisabled` | disabled                        |