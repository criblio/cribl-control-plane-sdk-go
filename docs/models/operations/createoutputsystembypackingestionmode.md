# CreateOutputSystemByPackIngestionMode

Ingestion mode

## Example Usage

```go
import (
	"github.com/Cribl-Community/cribl-control-plane-sdk-go/models/operations"
)

value := operations.CreateOutputSystemByPackIngestionModeBatching

// Open enum: custom values can be created with a direct type cast
custom := operations.CreateOutputSystemByPackIngestionMode("custom_value")
```


## Values

| Name                                             | Value                                            |
| ------------------------------------------------ | ------------------------------------------------ |
| `CreateOutputSystemByPackIngestionModeBatching`  | batching                                         |
| `CreateOutputSystemByPackIngestionModeStreaming` | streaming                                        |