# InputResponseInputWinEventLogsReadMode

Read all stored and future event logs, or only future events

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.InputResponseInputWinEventLogsReadModeOldest

// Open enum: custom values can be created with a direct type cast
custom := components.InputResponseInputWinEventLogsReadMode("custom_value")
```


## Values

| Name                                           | Value                                          |
| ---------------------------------------------- | ---------------------------------------------- |
| `InputResponseInputWinEventLogsReadModeOldest` | oldest                                         |
| `InputResponseInputWinEventLogsReadModeNewest` | newest                                         |