# InputWindowsMetricsCPUMode

Select the level of details for CPU metrics

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.InputWindowsMetricsCPUModeBasic

// Open enum: custom values can be created with a direct type cast
custom := components.InputWindowsMetricsCPUMode("custom_value")
```


## Values

| Name                                 | Value                                |
| ------------------------------------ | ------------------------------------ |
| `InputWindowsMetricsCPUModeBasic`    | basic                                |
| `InputWindowsMetricsCPUModeAll`      | all                                  |
| `InputWindowsMetricsCPUModeCustom`   | custom                               |
| `InputWindowsMetricsCPUModeDisabled` | disabled                             |