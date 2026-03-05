# IngestionMode

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.IngestionModeBatching

// Open enum: custom values can be created with a direct type cast
custom := components.IngestionMode("custom_value")
```


## Values

| Name                     | Value                    |
| ------------------------ | ------------------------ |
| `IngestionModeBatching`  | batching                 |
| `IngestionModeStreaming` | streaming                |