# WriteAction

Action to use when writing events. Must be set to `Create` when writing to a data stream.

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.WriteActionIndex

// Open enum: custom values can be created with a direct type cast
custom := components.WriteAction("custom_value")
```


## Values

| Name                | Value               |
| ------------------- | ------------------- |
| `WriteActionIndex`  | index               |
| `WriteActionCreate` | create              |