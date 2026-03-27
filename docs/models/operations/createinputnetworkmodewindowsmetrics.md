# CreateInputNetworkModeWindowsMetrics

Select the level of details for network metrics

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/operations"
)

value := operations.CreateInputNetworkModeWindowsMetricsBasic

// Open enum: custom values can be created with a direct type cast
custom := operations.CreateInputNetworkModeWindowsMetrics("custom_value")
```


## Values

| Name                                           | Value                                          |
| ---------------------------------------------- | ---------------------------------------------- |
| `CreateInputNetworkModeWindowsMetricsBasic`    | basic                                          |
| `CreateInputNetworkModeWindowsMetricsAll`      | all                                            |
| `CreateInputNetworkModeWindowsMetricsCustom`   | custom                                         |
| `CreateInputNetworkModeWindowsMetricsDisabled` | disabled                                       |