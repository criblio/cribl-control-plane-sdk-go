# MemoryModeWindowsMetrics

Select the level of details for memory metrics

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.MemoryModeWindowsMetricsBasic

// Open enum: custom values can be created with a direct type cast
custom := components.MemoryModeWindowsMetrics("custom_value")
```


## Values

| Name                               | Value                              |
| ---------------------------------- | ---------------------------------- |
| `MemoryModeWindowsMetricsBasic`    | basic                              |
| `MemoryModeWindowsMetricsAll`      | all                                |
| `MemoryModeWindowsMetricsCustom`   | custom                             |
| `MemoryModeWindowsMetricsDisabled` | disabled                           |