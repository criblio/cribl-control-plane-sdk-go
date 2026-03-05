# OutputSyslogSeverity

Default value for message severity. Will be overwritten by value of __severity if set. Defaults to notice.

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.OutputSyslogSeverityEmergency

// Open enum: custom values can be created with a direct type cast
custom := components.OutputSyslogSeverity(999)
```


## Values

| Name                            | Value                           |
| ------------------------------- | ------------------------------- |
| `OutputSyslogSeverityEmergency` | 0                               |
| `OutputSyslogSeverityAlert`     | 1                               |
| `OutputSyslogSeverityCritical`  | 2                               |
| `OutputSyslogSeverityError`     | 3                               |
| `OutputSyslogSeverityWarning`   | 4                               |
| `OutputSyslogSeverityNotice`    | 5                               |
| `OutputSyslogSeverityInfo`      | 6                               |
| `OutputSyslogSeverityDebug`     | 7                               |