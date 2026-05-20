# OutputResponseIngestionMode

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.OutputResponseIngestionModeBatching

// Open enum: custom values can be created with a direct type cast
custom := components.OutputResponseIngestionMode("custom_value")
```


## Values

| Name                                   | Value                                  |
| -------------------------------------- | -------------------------------------- |
| `OutputResponseIngestionModeBatching`  | batching                               |
| `OutputResponseIngestionModeStreaming` | streaming                              |