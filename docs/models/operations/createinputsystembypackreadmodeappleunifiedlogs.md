# CreateInputSystemByPackReadModeAppleUnifiedLogs

Read all log entries (historical and upcoming), or only upcoming, from the last entry

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/operations"
)

value := operations.CreateInputSystemByPackReadModeAppleUnifiedLogsOldest

// Open enum: custom values can be created with a direct type cast
custom := operations.CreateInputSystemByPackReadModeAppleUnifiedLogs("custom_value")
```


## Values

| Name                                                    | Value                                                   |
| ------------------------------------------------------- | ------------------------------------------------------- |
| `CreateInputSystemByPackReadModeAppleUnifiedLogsOldest` | oldest                                                  |
| `CreateInputSystemByPackReadModeAppleUnifiedLogsNewest` | newest                                                  |