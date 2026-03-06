# CreateInputSystemByPackCPUModeSystemMetrics

Select the level of detail for CPU metrics

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/operations"
)

value := operations.CreateInputSystemByPackCPUModeSystemMetricsBasic

// Open enum: custom values can be created with a direct type cast
custom := operations.CreateInputSystemByPackCPUModeSystemMetrics("custom_value")
```


## Values

| Name                                                  | Value                                                 |
| ----------------------------------------------------- | ----------------------------------------------------- |
| `CreateInputSystemByPackCPUModeSystemMetricsBasic`    | basic                                                 |
| `CreateInputSystemByPackCPUModeSystemMetricsAll`      | all                                                   |
| `CreateInputSystemByPackCPUModeSystemMetricsCustom`   | custom                                                |
| `CreateInputSystemByPackCPUModeSystemMetricsDisabled` | disabled                                              |