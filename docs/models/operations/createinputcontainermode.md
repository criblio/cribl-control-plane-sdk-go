# CreateInputContainerMode

Select the level of detail for container metrics

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/operations"
)

value := operations.CreateInputContainerModeBasic

// Open enum: custom values can be created with a direct type cast
custom := operations.CreateInputContainerMode("custom_value")
```


## Values

| Name                               | Value                              |
| ---------------------------------- | ---------------------------------- |
| `CreateInputContainerModeBasic`    | basic                              |
| `CreateInputContainerModeAll`      | all                                |
| `CreateInputContainerModeCustom`   | custom                             |
| `CreateInputContainerModeDisabled` | disabled                           |