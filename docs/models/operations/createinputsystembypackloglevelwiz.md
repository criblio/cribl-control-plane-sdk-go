# CreateInputSystemByPackLogLevelWiz

Collector runtime log level

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/operations"
)

value := operations.CreateInputSystemByPackLogLevelWizError

// Open enum: custom values can be created with a direct type cast
custom := operations.CreateInputSystemByPackLogLevelWiz("custom_value")
```


## Values

| Name                                      | Value                                     |
| ----------------------------------------- | ----------------------------------------- |
| `CreateInputSystemByPackLogLevelWizError` | error                                     |
| `CreateInputSystemByPackLogLevelWizWarn`  | warn                                      |
| `CreateInputSystemByPackLogLevelWizInfo`  | info                                      |
| `CreateInputSystemByPackLogLevelWizDebug` | debug                                     |
| `CreateInputSystemByPackLogLevelWizSilly` | silly                                     |