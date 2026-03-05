# CreateInputSystemByPackQueueType

The queue type used (or created)

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/operations"
)

value := operations.CreateInputSystemByPackQueueTypeStandard

// Open enum: custom values can be created with a direct type cast
custom := operations.CreateInputSystemByPackQueueType("custom_value")
```


## Values

| Name                                       | Value                                      |
| ------------------------------------------ | ------------------------------------------ |
| `CreateInputSystemByPackQueueTypeStandard` | standard                                   |
| `CreateInputSystemByPackQueueTypeFifo`     | fifo                                       |