# CreateOutputSystemByPackWriteAction

Action to use when writing events. Must be set to `Create` when writing to a data stream.

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/operations"
)

value := operations.CreateOutputSystemByPackWriteActionIndex

// Open enum: custom values can be created with a direct type cast
custom := operations.CreateOutputSystemByPackWriteAction("custom_value")
```


## Values

| Name                                        | Value                                       |
| ------------------------------------------- | ------------------------------------------- |
| `CreateOutputSystemByPackWriteActionIndex`  | index                                       |
| `CreateOutputSystemByPackWriteActionCreate` | create                                      |