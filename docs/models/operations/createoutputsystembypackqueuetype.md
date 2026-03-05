# CreateOutputSystemByPackQueueType

The queue type used (or created). Defaults to Standard.

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/operations"
)

value := operations.CreateOutputSystemByPackQueueTypeStandard

// Open enum: custom values can be created with a direct type cast
custom := operations.CreateOutputSystemByPackQueueType("custom_value")
```


## Values

| Name                                        | Value                                       |
| ------------------------------------------- | ------------------------------------------- |
| `CreateOutputSystemByPackQueueTypeStandard` | standard                                    |
| `CreateOutputSystemByPackQueueTypeFifo`     | fifo                                        |