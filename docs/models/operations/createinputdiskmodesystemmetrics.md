# CreateInputDiskModeSystemMetrics

Select the level of detail for disk metrics

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/operations"
)

value := operations.CreateInputDiskModeSystemMetricsBasic

// Open enum: custom values can be created with a direct type cast
custom := operations.CreateInputDiskModeSystemMetrics("custom_value")
```


## Values

| Name                                       | Value                                      |
| ------------------------------------------ | ------------------------------------------ |
| `CreateInputDiskModeSystemMetricsBasic`    | basic                                      |
| `CreateInputDiskModeSystemMetricsAll`      | all                                        |
| `CreateInputDiskModeSystemMetricsCustom`   | custom                                     |
| `CreateInputDiskModeSystemMetricsDisabled` | disabled                                   |