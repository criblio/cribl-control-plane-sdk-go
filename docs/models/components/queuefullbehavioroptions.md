# QueueFullBehaviorOptions

How to handle events when the queue is exerting backpressure (full capacity or low disk). 'Block' is the same behavior as non-PQ blocking. 'Drop new data' throws away incoming data, while leaving the contents of the PQ unchanged.

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.QueueFullBehaviorOptionsBlock

// Open enum: custom values can be created with a direct type cast
custom := components.QueueFullBehaviorOptions("custom_value")
```


## Values

| Name                            | Value                           |
| ------------------------------- | ------------------------------- |
| `QueueFullBehaviorOptionsBlock` | block                           |
| `QueueFullBehaviorOptionsDrop`  | drop                            |