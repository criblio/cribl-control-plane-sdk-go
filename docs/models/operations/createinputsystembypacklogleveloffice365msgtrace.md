# CreateInputSystemByPackLogLevelOffice365MsgTrace

Log Level (verbosity) for collection runtime behavior.

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/operations"
)

value := operations.CreateInputSystemByPackLogLevelOffice365MsgTraceError

// Open enum: custom values can be created with a direct type cast
custom := operations.CreateInputSystemByPackLogLevelOffice365MsgTrace("custom_value")
```


## Values

| Name                                                    | Value                                                   |
| ------------------------------------------------------- | ------------------------------------------------------- |
| `CreateInputSystemByPackLogLevelOffice365MsgTraceError` | error                                                   |
| `CreateInputSystemByPackLogLevelOffice365MsgTraceWarn`  | warn                                                    |
| `CreateInputSystemByPackLogLevelOffice365MsgTraceInfo`  | info                                                    |
| `CreateInputSystemByPackLogLevelOffice365MsgTraceDebug` | debug                                                   |
| `CreateInputSystemByPackLogLevelOffice365MsgTraceSilly` | silly                                                   |