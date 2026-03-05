# CreateInputReadMode

Read all stored and future event logs, or only future events

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/operations"
)

value := operations.CreateInputReadModeOldest

// Open enum: custom values can be created with a direct type cast
custom := operations.CreateInputReadMode("custom_value")
```


## Values

| Name                        | Value                       |
| --------------------------- | --------------------------- |
| `CreateInputReadModeOldest` | oldest                      |
| `CreateInputReadModeNewest` | newest                      |