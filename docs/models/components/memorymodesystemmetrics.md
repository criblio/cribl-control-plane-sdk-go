# MemoryModeSystemMetrics

Select the level of detail for memory metrics

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.MemoryModeSystemMetricsBasic

// Open enum: custom values can be created with a direct type cast
custom := components.MemoryModeSystemMetrics("custom_value")
```


## Values

| Name                              | Value                             |
| --------------------------------- | --------------------------------- |
| `MemoryModeSystemMetricsBasic`    | basic                             |
| `MemoryModeSystemMetricsAll`      | all                               |
| `MemoryModeSystemMetricsCustom`   | custom                            |
| `MemoryModeSystemMetricsDisabled` | disabled                          |