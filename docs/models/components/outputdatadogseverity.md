# OutputDatadogSeverity

Default value for message severity. When you send logs as JSON objects, the event's '__severity' field (if set) will override this value.

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.OutputDatadogSeverityEmergency

// Open enum: custom values can be created with a direct type cast
custom := components.OutputDatadogSeverity("custom_value")
```


## Values

| Name                             | Value                            |
| -------------------------------- | -------------------------------- |
| `OutputDatadogSeverityEmergency` | emergency                        |
| `OutputDatadogSeverityAlert`     | alert                            |
| `OutputDatadogSeverityCritical`  | critical                         |
| `OutputDatadogSeverityError`     | error                            |
| `OutputDatadogSeverityWarning`   | warning                          |
| `OutputDatadogSeverityNotice`    | notice                           |
| `OutputDatadogSeverityInfo`      | info                             |
| `OutputDatadogSeverityDebug`     | debug                            |