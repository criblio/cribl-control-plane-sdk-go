# NetworkModeSystemMetrics

Select the level of detail for network metrics

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.NetworkModeSystemMetricsBasic

// Open enum: custom values can be created with a direct type cast
custom := components.NetworkModeSystemMetrics("custom_value")
```


## Values

| Name                               | Value                              |
| ---------------------------------- | ---------------------------------- |
| `NetworkModeSystemMetricsBasic`    | basic                              |
| `NetworkModeSystemMetricsAll`      | all                                |
| `NetworkModeSystemMetricsCustom`   | custom                             |
| `NetworkModeSystemMetricsDisabled` | disabled                           |