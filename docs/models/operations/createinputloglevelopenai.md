# CreateInputLogLevelOpenai

Collector runtime log level.

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/operations"
)

value := operations.CreateInputLogLevelOpenaiError

// Open enum: custom values can be created with a direct type cast
custom := operations.CreateInputLogLevelOpenai("custom_value")
```


## Values

| Name                             | Value                            |
| -------------------------------- | -------------------------------- |
| `CreateInputLogLevelOpenaiError` | error                            |
| `CreateInputLogLevelOpenaiWarn`  | warn                             |
| `CreateInputLogLevelOpenaiInfo`  | info                             |
| `CreateInputLogLevelOpenaiDebug` | debug                            |
| `CreateInputLogLevelOpenaiSilly` | silly                            |