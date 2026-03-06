# InputWindowsMetricsDiskMode

Select the level of details for disk metrics

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.InputWindowsMetricsDiskModeBasic

// Open enum: custom values can be created with a direct type cast
custom := components.InputWindowsMetricsDiskMode("custom_value")
```


## Values

| Name                                  | Value                                 |
| ------------------------------------- | ------------------------------------- |
| `InputWindowsMetricsDiskModeBasic`    | basic                                 |
| `InputWindowsMetricsDiskModeAll`      | all                                   |
| `InputWindowsMetricsDiskModeCustom`   | custom                                |
| `InputWindowsMetricsDiskModeDisabled` | disabled                              |