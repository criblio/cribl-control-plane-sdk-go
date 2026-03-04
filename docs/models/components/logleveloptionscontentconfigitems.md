# LogLevelOptionsContentConfigItems

Collector runtime Log Level

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.LogLevelOptionsContentConfigItemsError

// Open enum: custom values can be created with a direct type cast
custom := components.LogLevelOptionsContentConfigItems("custom_value")
```


## Values

| Name                                     | Value                                    |
| ---------------------------------------- | ---------------------------------------- |
| `LogLevelOptionsContentConfigItemsError` | error                                    |
| `LogLevelOptionsContentConfigItemsWarn`  | warn                                     |
| `LogLevelOptionsContentConfigItemsInfo`  | info                                     |
| `LogLevelOptionsContentConfigItemsDebug` | debug                                    |