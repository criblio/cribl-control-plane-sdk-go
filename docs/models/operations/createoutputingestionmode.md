# CreateOutputIngestionMode

Ingestion mode

## Example Usage

```go
import (
	"github.com/Cribl-Community/cribl-control-plane-sdk-go/models/operations"
)

value := operations.CreateOutputIngestionModeBatching

// Open enum: custom values can be created with a direct type cast
custom := operations.CreateOutputIngestionMode("custom_value")
```


## Values

| Name                                 | Value                                |
| ------------------------------------ | ------------------------------------ |
| `CreateOutputIngestionModeBatching`  | batching                             |
| `CreateOutputIngestionModeStreaming` | streaming                            |