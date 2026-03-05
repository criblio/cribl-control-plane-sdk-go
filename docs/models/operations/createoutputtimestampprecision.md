# CreateOutputTimestampPrecision

Sets the precision for the supplied Unix time values. Defaults to milliseconds.

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/operations"
)

value := operations.CreateOutputTimestampPrecisionNs

// Open enum: custom values can be created with a direct type cast
custom := operations.CreateOutputTimestampPrecision("custom_value")
```


## Values

| Name                               | Value                              |
| ---------------------------------- | ---------------------------------- |
| `CreateOutputTimestampPrecisionNs` | ns                                 |
| `CreateOutputTimestampPrecisionU`  | u                                  |
| `CreateOutputTimestampPrecisionMs` | ms                                 |
| `CreateOutputTimestampPrecisionS`  | s                                  |
| `CreateOutputTimestampPrecisionM`  | m                                  |
| `CreateOutputTimestampPrecisionH`  | h                                  |