# BackpressureBehaviorOptions1

How to handle events when all receivers are exerting backpressure

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.BackpressureBehaviorOptions1Block

// Open enum: custom values can be created with a direct type cast
custom := components.BackpressureBehaviorOptions1("custom_value")
```


## Values

| Name                                | Value                               |
| ----------------------------------- | ----------------------------------- |
| `BackpressureBehaviorOptions1Block` | block                               |
| `BackpressureBehaviorOptions1Drop`  | drop                                |