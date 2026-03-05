# CreateOutputSystemByPackSeveritySyslog

Default value for message severity. Will be overwritten by value of __severity if set. Defaults to notice.

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/operations"
)

value := operations.CreateOutputSystemByPackSeveritySyslogEmergency

// Open enum: custom values can be created with a direct type cast
custom := operations.CreateOutputSystemByPackSeveritySyslog(999)
```


## Values

| Name                                              | Value                                             |
| ------------------------------------------------- | ------------------------------------------------- |
| `CreateOutputSystemByPackSeveritySyslogEmergency` | 0                                                 |
| `CreateOutputSystemByPackSeveritySyslogAlert`     | 1                                                 |
| `CreateOutputSystemByPackSeveritySyslogCritical`  | 2                                                 |
| `CreateOutputSystemByPackSeveritySyslogError`     | 3                                                 |
| `CreateOutputSystemByPackSeveritySyslogWarning`   | 4                                                 |
| `CreateOutputSystemByPackSeveritySyslogNotice`    | 5                                                 |
| `CreateOutputSystemByPackSeveritySyslogInfo`      | 6                                                 |
| `CreateOutputSystemByPackSeveritySyslogDebug`     | 7                                                 |