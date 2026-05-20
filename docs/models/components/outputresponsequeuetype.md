# OutputResponseQueueType

The queue type used (or created). Defaults to Standard.

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.OutputResponseQueueTypeStandard

// Open enum: custom values can be created with a direct type cast
custom := components.OutputResponseQueueType("custom_value")
```


## Values

| Name                              | Value                             |
| --------------------------------- | --------------------------------- |
| `OutputResponseQueueTypeStandard` | standard                          |
| `OutputResponseQueueTypeFifo`     | fifo                              |