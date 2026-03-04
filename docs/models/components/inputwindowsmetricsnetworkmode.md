# InputWindowsMetricsNetworkMode

Select the level of details for network metrics

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.InputWindowsMetricsNetworkModeBasic

// Open enum: custom values can be created with a direct type cast
custom := components.InputWindowsMetricsNetworkMode("custom_value")
```


## Values

| Name                                     | Value                                    |
| ---------------------------------------- | ---------------------------------------- |
| `InputWindowsMetricsNetworkModeBasic`    | basic                                    |
| `InputWindowsMetricsNetworkModeAll`      | all                                      |
| `InputWindowsMetricsNetworkModeCustom`   | custom                                   |
| `InputWindowsMetricsNetworkModeDisabled` | disabled                                 |