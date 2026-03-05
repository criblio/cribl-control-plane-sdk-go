# InputSystemMetricsMemoryMode

Select the level of detail for memory metrics

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.InputSystemMetricsMemoryModeBasic

// Open enum: custom values can be created with a direct type cast
custom := components.InputSystemMetricsMemoryMode("custom_value")
```


## Values

| Name                                   | Value                                  |
| -------------------------------------- | -------------------------------------- |
| `InputSystemMetricsMemoryModeBasic`    | basic                                  |
| `InputSystemMetricsMemoryModeAll`      | all                                    |
| `InputSystemMetricsMemoryModeCustom`   | custom                                 |
| `InputSystemMetricsMemoryModeDisabled` | disabled                               |