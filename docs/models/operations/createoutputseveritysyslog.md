# CreateOutputSeveritySyslog

Default value for message severity. Will be overwritten by value of __severity if set. Defaults to notice.

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/operations"
)

value := operations.CreateOutputSeveritySyslogEmergency

// Open enum: custom values can be created with a direct type cast
custom := operations.CreateOutputSeveritySyslog(999)
```


## Values

| Name                                  | Value                                 |
| ------------------------------------- | ------------------------------------- |
| `CreateOutputSeveritySyslogEmergency` | 0                                     |
| `CreateOutputSeveritySyslogAlert`     | 1                                     |
| `CreateOutputSeveritySyslogCritical`  | 2                                     |
| `CreateOutputSeveritySyslogError`     | 3                                     |
| `CreateOutputSeveritySyslogWarning`   | 4                                     |
| `CreateOutputSeveritySyslogNotice`    | 5                                     |
| `CreateOutputSeveritySyslogInfo`      | 6                                     |
| `CreateOutputSeveritySyslogDebug`     | 7                                     |