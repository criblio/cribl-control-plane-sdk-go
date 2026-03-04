# CreateInputNetworkModeSystemMetrics

Select the level of detail for network metrics

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/operations"
)

value := operations.CreateInputNetworkModeSystemMetricsBasic

// Open enum: custom values can be created with a direct type cast
custom := operations.CreateInputNetworkModeSystemMetrics("custom_value")
```


## Values

| Name                                          | Value                                         |
| --------------------------------------------- | --------------------------------------------- |
| `CreateInputNetworkModeSystemMetricsBasic`    | basic                                         |
| `CreateInputNetworkModeSystemMetricsAll`      | all                                           |
| `CreateInputNetworkModeSystemMetricsCustom`   | custom                                        |
| `CreateInputNetworkModeSystemMetricsDisabled` | disabled                                      |