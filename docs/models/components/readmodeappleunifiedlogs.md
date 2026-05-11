# ReadModeAppleUnifiedLogs

Read all log entries (historical and upcoming), or only upcoming, from the last entry

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.ReadModeAppleUnifiedLogsOldest

// Open enum: custom values can be created with a direct type cast
custom := components.ReadModeAppleUnifiedLogs("custom_value")
```


## Values

| Name                             | Value                            |
| -------------------------------- | -------------------------------- |
| `ReadModeAppleUnifiedLogsOldest` | oldest                           |
| `ReadModeAppleUnifiedLogsNewest` | newest                           |