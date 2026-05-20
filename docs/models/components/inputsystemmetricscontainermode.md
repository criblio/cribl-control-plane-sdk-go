# InputSystemMetricsContainerMode

Select the level of detail for container metrics

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.InputSystemMetricsContainerModeBasic

// Open enum: custom values can be created with a direct type cast
custom := components.InputSystemMetricsContainerMode("custom_value")
```


## Values

| Name                                      | Value                                     |
| ----------------------------------------- | ----------------------------------------- |
| `InputSystemMetricsContainerModeBasic`    | basic                                     |
| `InputSystemMetricsContainerModeAll`      | all                                       |
| `InputSystemMetricsContainerModeCustom`   | custom                                    |
| `InputSystemMetricsContainerModeDisabled` | disabled                                  |