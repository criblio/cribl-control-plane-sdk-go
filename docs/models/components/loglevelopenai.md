# LogLevelOpenai

Collector runtime log level.

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.LogLevelOpenaiError

// Open enum: custom values can be created with a direct type cast
custom := components.LogLevelOpenai("custom_value")
```


## Values

| Name                  | Value                 |
| --------------------- | --------------------- |
| `LogLevelOpenaiError` | error                 |
| `LogLevelOpenaiWarn`  | warn                  |
| `LogLevelOpenaiInfo`  | info                  |
| `LogLevelOpenaiDebug` | debug                 |
| `LogLevelOpenaiSilly` | silly                 |