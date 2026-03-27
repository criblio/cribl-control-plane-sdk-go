# InputWindowsMetricsMemoryMode

Select the level of details for memory metrics

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.InputWindowsMetricsMemoryModeBasic

// Open enum: custom values can be created with a direct type cast
custom := components.InputWindowsMetricsMemoryMode("custom_value")
```


## Values

| Name                                    | Value                                   |
| --------------------------------------- | --------------------------------------- |
| `InputWindowsMetricsMemoryModeBasic`    | basic                                   |
| `InputWindowsMetricsMemoryModeAll`      | all                                     |
| `InputWindowsMetricsMemoryModeCustom`   | custom                                  |
| `InputWindowsMetricsMemoryModeDisabled` | disabled                                |