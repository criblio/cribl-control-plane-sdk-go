# InputSqsQueueType

The queue type used (or created)

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.InputSqsQueueTypeStandard

// Open enum: custom values can be created with a direct type cast
custom := components.InputSqsQueueType("custom_value")
```


## Values

| Name                        | Value                       |
| --------------------------- | --------------------------- |
| `InputSqsQueueTypeStandard` | standard                    |
| `InputSqsQueueTypeFifo`     | fifo                        |