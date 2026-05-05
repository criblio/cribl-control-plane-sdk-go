# InputResponseContainerMode

Select the level of detail for container metrics

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.InputResponseContainerModeBasic

// Open enum: custom values can be created with a direct type cast
custom := components.InputResponseContainerMode("custom_value")
```


## Values

| Name                                 | Value                                |
| ------------------------------------ | ------------------------------------ |
| `InputResponseContainerModeBasic`    | basic                                |
| `InputResponseContainerModeAll`      | all                                  |
| `InputResponseContainerModeCustom`   | custom                               |
| `InputResponseContainerModeDisabled` | disabled                             |