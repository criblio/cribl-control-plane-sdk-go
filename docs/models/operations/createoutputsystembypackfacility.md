# CreateOutputSystemByPackFacility

Default value for message facility. Will be overwritten by value of __facility if set. Defaults to user.

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/operations"
)

value := operations.CreateOutputSystemByPackFacilityKern

// Open enum: custom values can be created with a direct type cast
custom := operations.CreateOutputSystemByPackFacility(999)
```


## Values

| Name                                          | Value                                         |
| --------------------------------------------- | --------------------------------------------- |
| `CreateOutputSystemByPackFacilityKern`        | 0                                             |
| `CreateOutputSystemByPackFacilityUser`        | 1                                             |
| `CreateOutputSystemByPackFacilityMail`        | 2                                             |
| `CreateOutputSystemByPackFacilityDaemon`      | 3                                             |
| `CreateOutputSystemByPackFacilityAuth`        | 4                                             |
| `CreateOutputSystemByPackFacilitySyslog`      | 5                                             |
| `CreateOutputSystemByPackFacilityLpr`         | 6                                             |
| `CreateOutputSystemByPackFacilityNews`        | 7                                             |
| `CreateOutputSystemByPackFacilityUucp`        | 8                                             |
| `CreateOutputSystemByPackFacilityCron`        | 9                                             |
| `CreateOutputSystemByPackFacilityAuthpriv`    | 10                                            |
| `CreateOutputSystemByPackFacilityFtp`         | 11                                            |
| `CreateOutputSystemByPackFacilityNtp`         | 12                                            |
| `CreateOutputSystemByPackFacilitySecurity`    | 13                                            |
| `CreateOutputSystemByPackFacilityConsole`     | 14                                            |
| `CreateOutputSystemByPackFacilitySolarisCron` | 15                                            |
| `CreateOutputSystemByPackFacilityLocal0`      | 16                                            |
| `CreateOutputSystemByPackFacilityLocal1`      | 17                                            |
| `CreateOutputSystemByPackFacilityLocal2`      | 18                                            |
| `CreateOutputSystemByPackFacilityLocal3`      | 19                                            |
| `CreateOutputSystemByPackFacilityLocal4`      | 20                                            |
| `CreateOutputSystemByPackFacilityLocal5`      | 21                                            |