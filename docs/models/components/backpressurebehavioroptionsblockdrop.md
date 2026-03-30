# BackpressureBehaviorOptionsBlockDrop

How to handle events when all receivers are exerting backpressure

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.BackpressureBehaviorOptionsBlockDropBlock

// Open enum: custom values can be created with a direct type cast
custom := components.BackpressureBehaviorOptionsBlockDrop("custom_value")
```


## Values

| Name                                        | Value                                       |
| ------------------------------------------- | ------------------------------------------- |
| `BackpressureBehaviorOptionsBlockDropBlock` | block                                       |
| `BackpressureBehaviorOptionsBlockDropDrop`  | drop                                        |