# CreateInputLogLevelWiz

Collector runtime log level

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/operations"
)

value := operations.CreateInputLogLevelWizError

// Open enum: custom values can be created with a direct type cast
custom := operations.CreateInputLogLevelWiz("custom_value")
```


## Values

| Name                          | Value                         |
| ----------------------------- | ----------------------------- |
| `CreateInputLogLevelWizError` | error                         |
| `CreateInputLogLevelWizWarn`  | warn                          |
| `CreateInputLogLevelWizInfo`  | info                          |
| `CreateInputLogLevelWizDebug` | debug                         |
| `CreateInputLogLevelWizSilly` | silly                         |