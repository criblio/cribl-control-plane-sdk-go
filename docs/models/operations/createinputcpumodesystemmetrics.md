# CreateInputCPUModeSystemMetrics

Select the level of detail for CPU metrics

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/operations"
)

value := operations.CreateInputCPUModeSystemMetricsBasic

// Open enum: custom values can be created with a direct type cast
custom := operations.CreateInputCPUModeSystemMetrics("custom_value")
```


## Values

| Name                                      | Value                                     |
| ----------------------------------------- | ----------------------------------------- |
| `CreateInputCPUModeSystemMetricsBasic`    | basic                                     |
| `CreateInputCPUModeSystemMetricsAll`      | all                                       |
| `CreateInputCPUModeSystemMetricsCustom`   | custom                                    |
| `CreateInputCPUModeSystemMetricsDisabled` | disabled                                  |