# CreateInputSystemModeSystemMetrics

Select the level of detail for system metrics

## Example Usage

```go
import (
	"github.com/Cribl-Community/cribl-control-plane-sdk-go/models/operations"
)

value := operations.CreateInputSystemModeSystemMetricsBasic

// Open enum: custom values can be created with a direct type cast
custom := operations.CreateInputSystemModeSystemMetrics("custom_value")
```


## Values

| Name                                         | Value                                        |
| -------------------------------------------- | -------------------------------------------- |
| `CreateInputSystemModeSystemMetricsBasic`    | basic                                        |
| `CreateInputSystemModeSystemMetricsAll`      | all                                          |
| `CreateInputSystemModeSystemMetricsCustom`   | custom                                       |
| `CreateInputSystemModeSystemMetricsDisabled` | disabled                                     |