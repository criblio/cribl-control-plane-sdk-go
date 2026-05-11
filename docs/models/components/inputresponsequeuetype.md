# InputResponseQueueType

The queue type used (or created)

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.InputResponseQueueTypeStandard

// Open enum: custom values can be created with a direct type cast
custom := components.InputResponseQueueType("custom_value")
```


## Values

| Name                             | Value                            |
| -------------------------------- | -------------------------------- |
| `InputResponseQueueTypeStandard` | standard                         |
| `InputResponseQueueTypeFifo`     | fifo                             |