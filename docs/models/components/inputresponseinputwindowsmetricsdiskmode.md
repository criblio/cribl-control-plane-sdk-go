# InputResponseInputWindowsMetricsDiskMode

Select the level of details for disk metrics

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.InputResponseInputWindowsMetricsDiskModeBasic

// Open enum: custom values can be created with a direct type cast
custom := components.InputResponseInputWindowsMetricsDiskMode("custom_value")
```


## Values

| Name                                               | Value                                              |
| -------------------------------------------------- | -------------------------------------------------- |
| `InputResponseInputWindowsMetricsDiskModeBasic`    | basic                                              |
| `InputResponseInputWindowsMetricsDiskModeAll`      | all                                                |
| `InputResponseInputWindowsMetricsDiskModeCustom`   | custom                                             |
| `InputResponseInputWindowsMetricsDiskModeDisabled` | disabled                                           |