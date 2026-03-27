# CreateOutputSystemByPackFormatDynatraceHTTP

How to format events before sending. Defaults to JSON. Plaintext is not currently supported.

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/operations"
)

value := operations.CreateOutputSystemByPackFormatDynatraceHTTPJSONArray

// Open enum: custom values can be created with a direct type cast
custom := operations.CreateOutputSystemByPackFormatDynatraceHTTP("custom_value")
```


## Values

| Name                                                   | Value                                                  |
| ------------------------------------------------------ | ------------------------------------------------------ |
| `CreateOutputSystemByPackFormatDynatraceHTTPJSONArray` | json_array                                             |
| `CreateOutputSystemByPackFormatDynatraceHTTPPlaintext` | plaintext                                              |