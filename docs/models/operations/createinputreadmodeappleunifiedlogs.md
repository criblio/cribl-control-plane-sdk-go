# CreateInputReadModeAppleUnifiedLogs

Read all log entries (historical and upcoming), or only upcoming, from the last entry

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/operations"
)

value := operations.CreateInputReadModeAppleUnifiedLogsOldest

// Open enum: custom values can be created with a direct type cast
custom := operations.CreateInputReadModeAppleUnifiedLogs("custom_value")
```


## Values

| Name                                        | Value                                       |
| ------------------------------------------- | ------------------------------------------- |
| `CreateInputReadModeAppleUnifiedLogsOldest` | oldest                                      |
| `CreateInputReadModeAppleUnifiedLogsNewest` | newest                                      |