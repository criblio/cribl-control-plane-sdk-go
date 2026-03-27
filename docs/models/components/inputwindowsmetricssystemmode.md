# InputWindowsMetricsSystemMode

Select the level of details for system metrics

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.InputWindowsMetricsSystemModeBasic

// Open enum: custom values can be created with a direct type cast
custom := components.InputWindowsMetricsSystemMode("custom_value")
```


## Values

| Name                                    | Value                                   |
| --------------------------------------- | --------------------------------------- |
| `InputWindowsMetricsSystemModeBasic`    | basic                                   |
| `InputWindowsMetricsSystemModeAll`      | all                                     |
| `InputWindowsMetricsSystemModeCustom`   | custom                                  |
| `InputWindowsMetricsSystemModeDisabled` | disabled                                |