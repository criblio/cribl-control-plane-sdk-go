# OutputResponseTimestampPrecision

Sets the precision for the supplied Unix time values. Defaults to milliseconds.

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.OutputResponseTimestampPrecisionNs

// Open enum: custom values can be created with a direct type cast
custom := components.OutputResponseTimestampPrecision("custom_value")
```


## Values

| Name                                 | Value                                |
| ------------------------------------ | ------------------------------------ |
| `OutputResponseTimestampPrecisionNs` | ns                                   |
| `OutputResponseTimestampPrecisionU`  | u                                    |
| `OutputResponseTimestampPrecisionMs` | ms                                   |
| `OutputResponseTimestampPrecisionS`  | s                                    |
| `OutputResponseTimestampPrecisionM`  | m                                    |
| `OutputResponseTimestampPrecisionH`  | h                                    |