# QueueFullBehaviorOptionsPq

Whether to block or drop events when the queue is exerting backpressure (full capacity or low disk). 'Block' is the same behavior as non-PQ blocking. 'Drop new data' throws away incoming data, while leaving the contents of the PQ unchanged.

## Example Usage

```go
import (
	"github.com/Cribl-Community/cribl-control-plane-sdk-go/models/components"
)

value := components.QueueFullBehaviorOptionsPqBlock

// Open enum: custom values can be created with a direct type cast
custom := components.QueueFullBehaviorOptionsPq("custom_value")
```


## Values

| Name                              | Value                             |
| --------------------------------- | --------------------------------- |
| `QueueFullBehaviorOptionsPqBlock` | block                             |
| `QueueFullBehaviorOptionsPqDrop`  | drop                              |