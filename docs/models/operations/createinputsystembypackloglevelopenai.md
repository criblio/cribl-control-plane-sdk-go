# CreateInputSystemByPackLogLevelOpenai

Collector runtime log level.

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/operations"
)

value := operations.CreateInputSystemByPackLogLevelOpenaiError

// Open enum: custom values can be created with a direct type cast
custom := operations.CreateInputSystemByPackLogLevelOpenai("custom_value")
```


## Values

| Name                                         | Value                                        |
| -------------------------------------------- | -------------------------------------------- |
| `CreateInputSystemByPackLogLevelOpenaiError` | error                                        |
| `CreateInputSystemByPackLogLevelOpenaiWarn`  | warn                                         |
| `CreateInputSystemByPackLogLevelOpenaiInfo`  | info                                         |
| `CreateInputSystemByPackLogLevelOpenaiDebug` | debug                                        |
| `CreateInputSystemByPackLogLevelOpenaiSilly` | silly                                        |