# CreateInputQueueType

The queue type used (or created)

## Example Usage

```go
import (
	"github.com/Cribl-Community/cribl-control-plane-sdk-go/models/operations"
)

value := operations.CreateInputQueueTypeStandard

// Open enum: custom values can be created with a direct type cast
custom := operations.CreateInputQueueType("custom_value")
```


## Values

| Name                           | Value                          |
| ------------------------------ | ------------------------------ |
| `CreateInputQueueTypeStandard` | standard                       |
| `CreateInputQueueTypeFifo`     | fifo                           |