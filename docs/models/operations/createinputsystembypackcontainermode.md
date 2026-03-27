# CreateInputSystemByPackContainerMode

Select the level of detail for container metrics

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/operations"
)

value := operations.CreateInputSystemByPackContainerModeBasic

// Open enum: custom values can be created with a direct type cast
custom := operations.CreateInputSystemByPackContainerMode("custom_value")
```


## Values

| Name                                           | Value                                          |
| ---------------------------------------------- | ---------------------------------------------- |
| `CreateInputSystemByPackContainerModeBasic`    | basic                                          |
| `CreateInputSystemByPackContainerModeAll`      | all                                            |
| `CreateInputSystemByPackContainerModeCustom`   | custom                                         |
| `CreateInputSystemByPackContainerModeDisabled` | disabled                                       |