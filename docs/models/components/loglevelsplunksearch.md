# LogLevelSplunkSearch

Collector runtime log level (verbosity)

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.LogLevelSplunkSearchError

// Open enum: custom values can be created with a direct type cast
custom := components.LogLevelSplunkSearch("custom_value")
```


## Values

| Name                        | Value                       |
| --------------------------- | --------------------------- |
| `LogLevelSplunkSearchError` | error                       |
| `LogLevelSplunkSearchWarn`  | warn                        |
| `LogLevelSplunkSearchInfo`  | info                        |
| `LogLevelSplunkSearchDebug` | debug                       |