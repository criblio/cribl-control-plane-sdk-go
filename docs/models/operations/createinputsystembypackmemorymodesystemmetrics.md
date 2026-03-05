# CreateInputSystemByPackMemoryModeSystemMetrics

Select the level of detail for memory metrics

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/operations"
)

value := operations.CreateInputSystemByPackMemoryModeSystemMetricsBasic

// Open enum: custom values can be created with a direct type cast
custom := operations.CreateInputSystemByPackMemoryModeSystemMetrics("custom_value")
```


## Values

| Name                                                     | Value                                                    |
| -------------------------------------------------------- | -------------------------------------------------------- |
| `CreateInputSystemByPackMemoryModeSystemMetricsBasic`    | basic                                                    |
| `CreateInputSystemByPackMemoryModeSystemMetricsAll`      | all                                                      |
| `CreateInputSystemByPackMemoryModeSystemMetricsCustom`   | custom                                                   |
| `CreateInputSystemByPackMemoryModeSystemMetricsDisabled` | disabled                                                 |