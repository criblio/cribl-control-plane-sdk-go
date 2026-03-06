# LogLevelForFailedLookups

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.LogLevelForFailedLookupsSilly

// Open enum: custom values can be created with a direct type cast
custom := components.LogLevelForFailedLookups("custom_value")
```


## Values

| Name                            | Value                           |
| ------------------------------- | ------------------------------- |
| `LogLevelForFailedLookupsSilly` | silly                           |
| `LogLevelForFailedLookupsDebug` | debug                           |
| `LogLevelForFailedLookupsInfo`  | info                            |
| `LogLevelForFailedLookupsWarn`  | warn                            |
| `LogLevelForFailedLookupsError` | error                           |