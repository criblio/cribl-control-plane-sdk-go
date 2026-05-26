# CreateOutputFacility

Default value for message facility. Will be overwritten by value of __facility if set. Defaults to user.

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/operations"
)

value := operations.CreateOutputFacilityKern

// Open enum: custom values can be created with a direct type cast
custom := operations.CreateOutputFacility(999)
```


## Values

| Name                              | Value                             |
| --------------------------------- | --------------------------------- |
| `CreateOutputFacilityKern`        | 0                                 |
| `CreateOutputFacilityUser`        | 1                                 |
| `CreateOutputFacilityMail`        | 2                                 |
| `CreateOutputFacilityDaemon`      | 3                                 |
| `CreateOutputFacilityAuth`        | 4                                 |
| `CreateOutputFacilitySyslog`      | 5                                 |
| `CreateOutputFacilityLpr`         | 6                                 |
| `CreateOutputFacilityNews`        | 7                                 |
| `CreateOutputFacilityUucp`        | 8                                 |
| `CreateOutputFacilityCron`        | 9                                 |
| `CreateOutputFacilityAuthpriv`    | 10                                |
| `CreateOutputFacilityFtp`         | 11                                |
| `CreateOutputFacilityNtp`         | 12                                |
| `CreateOutputFacilitySecurity`    | 13                                |
| `CreateOutputFacilityConsole`     | 14                                |
| `CreateOutputFacilitySolarisCron` | 15                                |
| `CreateOutputFacilityLocal0`      | 16                                |
| `CreateOutputFacilityLocal1`      | 17                                |
| `CreateOutputFacilityLocal2`      | 18                                |
| `CreateOutputFacilityLocal3`      | 19                                |
| `CreateOutputFacilityLocal4`      | 20                                |
| `CreateOutputFacilityLocal5`      | 21                                |