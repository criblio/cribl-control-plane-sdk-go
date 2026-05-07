# OutputSyslogFacility

Default value for message facility. Will be overwritten by value of __facility if set. Defaults to user.

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.OutputSyslogFacilityKern

// Open enum: custom values can be created with a direct type cast
custom := components.OutputSyslogFacility(999)
```


## Values

| Name                              | Value                             |
| --------------------------------- | --------------------------------- |
| `OutputSyslogFacilityKern`        | 0                                 |
| `OutputSyslogFacilityUser`        | 1                                 |
| `OutputSyslogFacilityMail`        | 2                                 |
| `OutputSyslogFacilityDaemon`      | 3                                 |
| `OutputSyslogFacilityAuth`        | 4                                 |
| `OutputSyslogFacilitySyslog`      | 5                                 |
| `OutputSyslogFacilityLpr`         | 6                                 |
| `OutputSyslogFacilityNews`        | 7                                 |
| `OutputSyslogFacilityUucp`        | 8                                 |
| `OutputSyslogFacilityCron`        | 9                                 |
| `OutputSyslogFacilityAuthpriv`    | 10                                |
| `OutputSyslogFacilityFtp`         | 11                                |
| `OutputSyslogFacilityNtp`         | 12                                |
| `OutputSyslogFacilitySecurity`    | 13                                |
| `OutputSyslogFacilityConsole`     | 14                                |
| `OutputSyslogFacilitySolarisCron` | 15                                |
| `OutputSyslogFacilityLocal0`      | 16                                |
| `OutputSyslogFacilityLocal1`      | 17                                |
| `OutputSyslogFacilityLocal2`      | 18                                |
| `OutputSyslogFacilityLocal3`      | 19                                |
| `OutputSyslogFacilityLocal4`      | 20                                |
| `OutputSyslogFacilityLocal5`      | 21                                |