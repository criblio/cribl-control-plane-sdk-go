# InputSystemMetricsNetworkMode

Select the level of detail for network metrics

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.InputSystemMetricsNetworkModeBasic

// Open enum: custom values can be created with a direct type cast
custom := components.InputSystemMetricsNetworkMode("custom_value")
```


## Values

| Name                                    | Value                                   |
| --------------------------------------- | --------------------------------------- |
| `InputSystemMetricsNetworkModeBasic`    | basic                                   |
| `InputSystemMetricsNetworkModeAll`      | all                                     |
| `InputSystemMetricsNetworkModeCustom`   | custom                                  |
| `InputSystemMetricsNetworkModeDisabled` | disabled                                |