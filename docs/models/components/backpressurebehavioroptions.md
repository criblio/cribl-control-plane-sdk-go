# BackpressureBehaviorOptions

How to handle events when all receivers are exerting backpressure

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.BackpressureBehaviorOptionsBlock

// Open enum: custom values can be created with a direct type cast
custom := components.BackpressureBehaviorOptions("custom_value")
```


## Values

| Name                               | Value                              |
| ---------------------------------- | ---------------------------------- |
| `BackpressureBehaviorOptionsBlock` | block                              |
| `BackpressureBehaviorOptionsDrop`  | drop                               |
| `BackpressureBehaviorOptionsQueue` | queue                              |