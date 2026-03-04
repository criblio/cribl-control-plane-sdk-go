# CreateInputSystemByPackLogLevelSplunkSearch

Collector runtime log level (verbosity)

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/operations"
)

value := operations.CreateInputSystemByPackLogLevelSplunkSearchError

// Open enum: custom values can be created with a direct type cast
custom := operations.CreateInputSystemByPackLogLevelSplunkSearch("custom_value")
```


## Values

| Name                                               | Value                                              |
| -------------------------------------------------- | -------------------------------------------------- |
| `CreateInputSystemByPackLogLevelSplunkSearchError` | error                                              |
| `CreateInputSystemByPackLogLevelSplunkSearchWarn`  | warn                                               |
| `CreateInputSystemByPackLogLevelSplunkSearchInfo`  | info                                               |
| `CreateInputSystemByPackLogLevelSplunkSearchDebug` | debug                                              |