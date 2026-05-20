# ModeOptionsGpu

Select the level of detail for GPU metrics

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.ModeOptionsGpuBasic

// Open enum: custom values can be created with a direct type cast
custom := components.ModeOptionsGpu("custom_value")
```


## Values

| Name                     | Value                    |
| ------------------------ | ------------------------ |
| `ModeOptionsGpuBasic`    | basic                    |
| `ModeOptionsGpuAll`      | all                      |
| `ModeOptionsGpuCustom`   | custom                   |
| `ModeOptionsGpuDisabled` | disabled                 |