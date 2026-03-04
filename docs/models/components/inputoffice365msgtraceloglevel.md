# InputOffice365MsgTraceLogLevel

Log Level (verbosity) for collection runtime behavior.

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.InputOffice365MsgTraceLogLevelError

// Open enum: custom values can be created with a direct type cast
custom := components.InputOffice365MsgTraceLogLevel("custom_value")
```


## Values

| Name                                  | Value                                 |
| ------------------------------------- | ------------------------------------- |
| `InputOffice365MsgTraceLogLevelError` | error                                 |
| `InputOffice365MsgTraceLogLevelWarn`  | warn                                  |
| `InputOffice365MsgTraceLogLevelInfo`  | info                                  |
| `InputOffice365MsgTraceLogLevelDebug` | debug                                 |
| `InputOffice365MsgTraceLogLevelSilly` | silly                                 |