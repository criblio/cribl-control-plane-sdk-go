# SystemModeSystemMetrics

Select the level of detail for system metrics

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.SystemModeSystemMetricsBasic

// Open enum: custom values can be created with a direct type cast
custom := components.SystemModeSystemMetrics("custom_value")
```


## Values

| Name                              | Value                             |
| --------------------------------- | --------------------------------- |
| `SystemModeSystemMetricsBasic`    | basic                             |
| `SystemModeSystemMetricsAll`      | all                               |
| `SystemModeSystemMetricsCustom`   | custom                            |
| `SystemModeSystemMetricsDisabled` | disabled                          |