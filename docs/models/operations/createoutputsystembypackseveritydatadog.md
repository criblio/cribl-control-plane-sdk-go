# CreateOutputSystemByPackSeverityDatadog

Default value for message severity. When you send logs as JSON objects, the event's '__severity' field (if set) will override this value.

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/operations"
)

value := operations.CreateOutputSystemByPackSeverityDatadogEmergency

// Open enum: custom values can be created with a direct type cast
custom := operations.CreateOutputSystemByPackSeverityDatadog("custom_value")
```


## Values

| Name                                               | Value                                              |
| -------------------------------------------------- | -------------------------------------------------- |
| `CreateOutputSystemByPackSeverityDatadogEmergency` | emergency                                          |
| `CreateOutputSystemByPackSeverityDatadogAlert`     | alert                                              |
| `CreateOutputSystemByPackSeverityDatadogCritical`  | critical                                           |
| `CreateOutputSystemByPackSeverityDatadogError`     | error                                              |
| `CreateOutputSystemByPackSeverityDatadogWarning`   | warning                                            |
| `CreateOutputSystemByPackSeverityDatadogNotice`    | notice                                             |
| `CreateOutputSystemByPackSeverityDatadogInfo`      | info                                               |
| `CreateOutputSystemByPackSeverityDatadogDebug`     | debug                                              |