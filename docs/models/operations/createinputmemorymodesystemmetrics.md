# CreateInputMemoryModeSystemMetrics

Select the level of detail for memory metrics

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/operations"
)

value := operations.CreateInputMemoryModeSystemMetricsBasic

// Open enum: custom values can be created with a direct type cast
custom := operations.CreateInputMemoryModeSystemMetrics("custom_value")
```


## Values

| Name                                         | Value                                        |
| -------------------------------------------- | -------------------------------------------- |
| `CreateInputMemoryModeSystemMetricsBasic`    | basic                                        |
| `CreateInputMemoryModeSystemMetricsAll`      | all                                          |
| `CreateInputMemoryModeSystemMetricsCustom`   | custom                                       |
| `CreateInputMemoryModeSystemMetricsDisabled` | disabled                                     |