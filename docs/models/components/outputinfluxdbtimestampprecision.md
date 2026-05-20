# OutputInfluxdbTimestampPrecision

Sets the precision for the supplied Unix time values. Defaults to milliseconds.

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.OutputInfluxdbTimestampPrecisionNs

// Open enum: custom values can be created with a direct type cast
custom := components.OutputInfluxdbTimestampPrecision("custom_value")
```


## Values

| Name                                 | Value                                |
| ------------------------------------ | ------------------------------------ |
| `OutputInfluxdbTimestampPrecisionNs` | ns                                   |
| `OutputInfluxdbTimestampPrecisionU`  | u                                    |
| `OutputInfluxdbTimestampPrecisionMs` | ms                                   |
| `OutputInfluxdbTimestampPrecisionS`  | s                                    |
| `OutputInfluxdbTimestampPrecisionM`  | m                                    |
| `OutputInfluxdbTimestampPrecisionH`  | h                                    |