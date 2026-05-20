# OutputResponseFacility

Default value for message facility. Will be overwritten by value of __facility if set. Defaults to user.

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.OutputResponseFacilityKern

// Open enum: custom values can be created with a direct type cast
custom := components.OutputResponseFacility(999)
```


## Values

| Name                                | Value                               |
| ----------------------------------- | ----------------------------------- |
| `OutputResponseFacilityKern`        | 0                                   |
| `OutputResponseFacilityUser`        | 1                                   |
| `OutputResponseFacilityMail`        | 2                                   |
| `OutputResponseFacilityDaemon`      | 3                                   |
| `OutputResponseFacilityAuth`        | 4                                   |
| `OutputResponseFacilitySyslog`      | 5                                   |
| `OutputResponseFacilityLpr`         | 6                                   |
| `OutputResponseFacilityNews`        | 7                                   |
| `OutputResponseFacilityUucp`        | 8                                   |
| `OutputResponseFacilityCron`        | 9                                   |
| `OutputResponseFacilityAuthpriv`    | 10                                  |
| `OutputResponseFacilityFtp`         | 11                                  |
| `OutputResponseFacilityNtp`         | 12                                  |
| `OutputResponseFacilitySecurity`    | 13                                  |
| `OutputResponseFacilityConsole`     | 14                                  |
| `OutputResponseFacilitySolarisCron` | 15                                  |
| `OutputResponseFacilityLocal0`      | 16                                  |
| `OutputResponseFacilityLocal1`      | 17                                  |
| `OutputResponseFacilityLocal2`      | 18                                  |
| `OutputResponseFacilityLocal3`      | 19                                  |
| `OutputResponseFacilityLocal4`      | 20                                  |
| `OutputResponseFacilityLocal5`      | 21                                  |