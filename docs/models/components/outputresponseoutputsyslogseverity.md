# OutputResponseOutputSyslogSeverity

Default value for message severity. Will be overwritten by value of __severity if set. Defaults to notice.

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.OutputResponseOutputSyslogSeverityEmergency

// Open enum: custom values can be created with a direct type cast
custom := components.OutputResponseOutputSyslogSeverity(999)
```


## Values

| Name                                          | Value                                         |
| --------------------------------------------- | --------------------------------------------- |
| `OutputResponseOutputSyslogSeverityEmergency` | 0                                             |
| `OutputResponseOutputSyslogSeverityAlert`     | 1                                             |
| `OutputResponseOutputSyslogSeverityCritical`  | 2                                             |
| `OutputResponseOutputSyslogSeverityError`     | 3                                             |
| `OutputResponseOutputSyslogSeverityWarning`   | 4                                             |
| `OutputResponseOutputSyslogSeverityNotice`    | 5                                             |
| `OutputResponseOutputSyslogSeverityInfo`      | 6                                             |
| `OutputResponseOutputSyslogSeverityDebug`     | 7                                             |