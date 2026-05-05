# OutputResponseOutputDynatraceHTTPFormat

How to format events before sending. Defaults to JSON. Plaintext is not currently supported.

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.OutputResponseOutputDynatraceHTTPFormatJSONArray

// Open enum: custom values can be created with a direct type cast
custom := components.OutputResponseOutputDynatraceHTTPFormat("custom_value")
```


## Values

| Name                                               | Value                                              |
| -------------------------------------------------- | -------------------------------------------------- |
| `OutputResponseOutputDynatraceHTTPFormatJSONArray` | json_array                                         |
| `OutputResponseOutputDynatraceHTTPFormatPlaintext` | plaintext                                          |