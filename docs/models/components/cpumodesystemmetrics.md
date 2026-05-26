# CPUModeSystemMetrics

Select the level of detail for CPU metrics

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.CPUModeSystemMetricsBasic

// Open enum: custom values can be created with a direct type cast
custom := components.CPUModeSystemMetrics("custom_value")
```


## Values

| Name                           | Value                          |
| ------------------------------ | ------------------------------ |
| `CPUModeSystemMetricsBasic`    | basic                          |
| `CPUModeSystemMetricsAll`      | all                            |
| `CPUModeSystemMetricsCustom`   | custom                         |
| `CPUModeSystemMetricsDisabled` | disabled                       |