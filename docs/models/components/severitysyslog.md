# SeveritySyslog

Default value for message severity. Will be overwritten by value of __severity if set. Defaults to notice.

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.SeveritySyslogEmergency

// Open enum: custom values can be created with a direct type cast
custom := components.SeveritySyslog(999)
```


## Values

| Name                      | Value                     |
| ------------------------- | ------------------------- |
| `SeveritySyslogEmergency` | 0                         |
| `SeveritySyslogAlert`     | 1                         |
| `SeveritySyslogCritical`  | 2                         |
| `SeveritySyslogError`     | 3                         |
| `SeveritySyslogWarning`   | 4                         |
| `SeveritySyslogNotice`    | 5                         |
| `SeveritySyslogInfo`      | 6                         |
| `SeveritySyslogDebug`     | 7                         |