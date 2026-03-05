# CreateOutputSystemByPackTimestampPrecision

Sets the precision for the supplied Unix time values. Defaults to milliseconds.

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/operations"
)

value := operations.CreateOutputSystemByPackTimestampPrecisionNs

// Open enum: custom values can be created with a direct type cast
custom := operations.CreateOutputSystemByPackTimestampPrecision("custom_value")
```


## Values

| Name                                           | Value                                          |
| ---------------------------------------------- | ---------------------------------------------- |
| `CreateOutputSystemByPackTimestampPrecisionNs` | ns                                             |
| `CreateOutputSystemByPackTimestampPrecisionU`  | u                                              |
| `CreateOutputSystemByPackTimestampPrecisionMs` | ms                                             |
| `CreateOutputSystemByPackTimestampPrecisionS`  | s                                              |
| `CreateOutputSystemByPackTimestampPrecisionM`  | m                                              |
| `CreateOutputSystemByPackTimestampPrecisionH`  | h                                              |