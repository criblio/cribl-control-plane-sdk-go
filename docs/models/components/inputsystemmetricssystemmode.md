# InputSystemMetricsSystemMode

Select the level of detail for system metrics

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.InputSystemMetricsSystemModeBasic

// Open enum: custom values can be created with a direct type cast
custom := components.InputSystemMetricsSystemMode("custom_value")
```


## Values

| Name                                   | Value                                  |
| -------------------------------------- | -------------------------------------- |
| `InputSystemMetricsSystemModeBasic`    | basic                                  |
| `InputSystemMetricsSystemModeAll`      | all                                    |
| `InputSystemMetricsSystemModeCustom`   | custom                                 |
| `InputSystemMetricsSystemModeDisabled` | disabled                               |