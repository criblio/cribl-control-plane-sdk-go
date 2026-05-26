# OutputResponseWriteAction

Action to use when writing events. Must be set to `Create` when writing to a data stream.

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.OutputResponseWriteActionIndex

// Open enum: custom values can be created with a direct type cast
custom := components.OutputResponseWriteAction("custom_value")
```


## Values

| Name                              | Value                             |
| --------------------------------- | --------------------------------- |
| `OutputResponseWriteActionIndex`  | index                             |
| `OutputResponseWriteActionCreate` | create                            |