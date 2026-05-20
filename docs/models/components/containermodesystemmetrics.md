# ContainerModeSystemMetrics

Select the level of detail for container metrics

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.ContainerModeSystemMetricsBasic

// Open enum: custom values can be created with a direct type cast
custom := components.ContainerModeSystemMetrics("custom_value")
```


## Values

| Name                                 | Value                                |
| ------------------------------------ | ------------------------------------ |
| `ContainerModeSystemMetricsBasic`    | basic                                |
| `ContainerModeSystemMetricsAll`      | all                                  |
| `ContainerModeSystemMetricsCustom`   | custom                               |
| `ContainerModeSystemMetricsDisabled` | disabled                             |