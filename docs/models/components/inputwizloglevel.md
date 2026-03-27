# InputWizLogLevel

Collector runtime log level

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.InputWizLogLevelError

// Open enum: custom values can be created with a direct type cast
custom := components.InputWizLogLevel("custom_value")
```


## Values

| Name                    | Value                   |
| ----------------------- | ----------------------- |
| `InputWizLogLevelError` | error                   |
| `InputWizLogLevelWarn`  | warn                    |
| `InputWizLogLevelInfo`  | info                    |
| `InputWizLogLevelDebug` | debug                   |
| `InputWizLogLevelSilly` | silly                   |