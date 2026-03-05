# ContainerMode

Select the level of detail for container metrics

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.ContainerModeBasic

// Open enum: custom values can be created with a direct type cast
custom := components.ContainerMode("custom_value")
```


## Values

| Name                    | Value                   |
| ----------------------- | ----------------------- |
| `ContainerModeBasic`    | basic                   |
| `ContainerModeAll`      | all                     |
| `ContainerModeCustom`   | custom                  |
| `ContainerModeDisabled` | disabled                |