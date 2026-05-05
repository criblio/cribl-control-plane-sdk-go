# InputWinEventLogsReadMode

Read all stored and future event logs, or only future events

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.InputWinEventLogsReadModeOldest

// Open enum: custom values can be created with a direct type cast
custom := components.InputWinEventLogsReadMode("custom_value")
```


## Values

| Name                              | Value                             |
| --------------------------------- | --------------------------------- |
| `InputWinEventLogsReadModeOldest` | oldest                            |
| `InputWinEventLogsReadModeNewest` | newest                            |