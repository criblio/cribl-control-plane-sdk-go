# CreateInputLogLevelOffice365MsgTrace

Log Level (verbosity) for collection runtime behavior.

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/operations"
)

value := operations.CreateInputLogLevelOffice365MsgTraceError

// Open enum: custom values can be created with a direct type cast
custom := operations.CreateInputLogLevelOffice365MsgTrace("custom_value")
```


## Values

| Name                                        | Value                                       |
| ------------------------------------------- | ------------------------------------------- |
| `CreateInputLogLevelOffice365MsgTraceError` | error                                       |
| `CreateInputLogLevelOffice365MsgTraceWarn`  | warn                                        |
| `CreateInputLogLevelOffice365MsgTraceInfo`  | info                                        |
| `CreateInputLogLevelOffice365MsgTraceDebug` | debug                                       |
| `CreateInputLogLevelOffice365MsgTraceSilly` | silly                                       |