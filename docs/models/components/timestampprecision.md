# TimestampPrecision

Sets the precision for the supplied Unix time values. Defaults to milliseconds.

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.TimestampPrecisionNs

// Open enum: custom values can be created with a direct type cast
custom := components.TimestampPrecision("custom_value")
```


## Values

| Name                   | Value                  |
| ---------------------- | ---------------------- |
| `TimestampPrecisionNs` | ns                     |
| `TimestampPrecisionU`  | u                      |
| `TimestampPrecisionMs` | ms                     |
| `TimestampPrecisionS`  | s                      |
| `TimestampPrecisionM`  | m                      |
| `TimestampPrecisionH`  | h                      |