# CreateInputLogLevelSplunkSearch

Collector runtime log level (verbosity)

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/operations"
)

value := operations.CreateInputLogLevelSplunkSearchError

// Open enum: custom values can be created with a direct type cast
custom := operations.CreateInputLogLevelSplunkSearch("custom_value")
```


## Values

| Name                                   | Value                                  |
| -------------------------------------- | -------------------------------------- |
| `CreateInputLogLevelSplunkSearchError` | error                                  |
| `CreateInputLogLevelSplunkSearchWarn`  | warn                                   |
| `CreateInputLogLevelSplunkSearchInfo`  | info                                   |
| `CreateInputLogLevelSplunkSearchDebug` | debug                                  |