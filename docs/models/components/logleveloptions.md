# LogLevelOptions

Collector runtime log level

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.LogLevelOptionsError

// Open enum: custom values can be created with a direct type cast
custom := components.LogLevelOptions("custom_value")
```


## Values

| Name                   | Value                  |
| ---------------------- | ---------------------- |
| `LogLevelOptionsError` | error                  |
| `LogLevelOptionsWarn`  | warn                   |
| `LogLevelOptionsInfo`  | info                   |
| `LogLevelOptionsDebug` | debug                  |