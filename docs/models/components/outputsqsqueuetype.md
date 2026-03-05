# OutputSqsQueueType

The queue type used (or created). Defaults to Standard.

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.OutputSqsQueueTypeStandard

// Open enum: custom values can be created with a direct type cast
custom := components.OutputSqsQueueType("custom_value")
```


## Values

| Name                         | Value                        |
| ---------------------------- | ---------------------------- |
| `OutputSqsQueueTypeStandard` | standard                     |
| `OutputSqsQueueTypeFifo`     | fifo                         |