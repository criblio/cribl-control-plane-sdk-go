# CreateOutputFormatDynatraceHTTP

How to format events before sending. Defaults to JSON. Plaintext is not currently supported.

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/operations"
)

value := operations.CreateOutputFormatDynatraceHTTPJSONArray

// Open enum: custom values can be created with a direct type cast
custom := operations.CreateOutputFormatDynatraceHTTP("custom_value")
```


## Values

| Name                                       | Value                                      |
| ------------------------------------------ | ------------------------------------------ |
| `CreateOutputFormatDynatraceHTTPJSONArray` | json_array                                 |
| `CreateOutputFormatDynatraceHTTPPlaintext` | plaintext                                  |