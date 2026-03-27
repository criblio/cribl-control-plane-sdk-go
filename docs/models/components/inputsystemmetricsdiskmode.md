# InputSystemMetricsDiskMode

Select the level of detail for disk metrics

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.InputSystemMetricsDiskModeBasic

// Open enum: custom values can be created with a direct type cast
custom := components.InputSystemMetricsDiskMode("custom_value")
```


## Values

| Name                                 | Value                                |
| ------------------------------------ | ------------------------------------ |
| `InputSystemMetricsDiskModeBasic`    | basic                                |
| `InputSystemMetricsDiskModeAll`      | all                                  |
| `InputSystemMetricsDiskModeCustom`   | custom                               |
| `InputSystemMetricsDiskModeDisabled` | disabled                             |