# CreateOutputSeverityDatadog

Default value for message severity. When you send logs as JSON objects, the event's '__severity' field (if set) will override this value.

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/operations"
)

value := operations.CreateOutputSeverityDatadogEmergency

// Open enum: custom values can be created with a direct type cast
custom := operations.CreateOutputSeverityDatadog("custom_value")
```


## Values

| Name                                   | Value                                  |
| -------------------------------------- | -------------------------------------- |
| `CreateOutputSeverityDatadogEmergency` | emergency                              |
| `CreateOutputSeverityDatadogAlert`     | alert                                  |
| `CreateOutputSeverityDatadogCritical`  | critical                               |
| `CreateOutputSeverityDatadogError`     | error                                  |
| `CreateOutputSeverityDatadogWarning`   | warning                                |
| `CreateOutputSeverityDatadogNotice`    | notice                                 |
| `CreateOutputSeverityDatadogInfo`      | info                                   |
| `CreateOutputSeverityDatadogDebug`     | debug                                  |