# CreateInputSystemByPackNetworkModeSystemMetrics

Select the level of detail for network metrics

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/operations"
)

value := operations.CreateInputSystemByPackNetworkModeSystemMetricsBasic

// Open enum: custom values can be created with a direct type cast
custom := operations.CreateInputSystemByPackNetworkModeSystemMetrics("custom_value")
```


## Values

| Name                                                      | Value                                                     |
| --------------------------------------------------------- | --------------------------------------------------------- |
| `CreateInputSystemByPackNetworkModeSystemMetricsBasic`    | basic                                                     |
| `CreateInputSystemByPackNetworkModeSystemMetricsAll`      | all                                                       |
| `CreateInputSystemByPackNetworkModeSystemMetricsCustom`   | custom                                                    |
| `CreateInputSystemByPackNetworkModeSystemMetricsDisabled` | disabled                                                  |