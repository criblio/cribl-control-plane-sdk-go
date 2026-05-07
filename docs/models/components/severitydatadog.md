# SeverityDatadog

Default value for message severity. When you send logs as JSON objects, the event's '__severity' field (if set) will override this value.

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.SeverityDatadogEmergency

// Open enum: custom values can be created with a direct type cast
custom := components.SeverityDatadog("custom_value")
```


## Values

| Name                       | Value                      |
| -------------------------- | -------------------------- |
| `SeverityDatadogEmergency` | emergency                  |
| `SeverityDatadogAlert`     | alert                      |
| `SeverityDatadogCritical`  | critical                   |
| `SeverityDatadogError`     | error                      |
| `SeverityDatadogWarning`   | warning                    |
| `SeverityDatadogNotice`    | notice                     |
| `SeverityDatadogInfo`      | info                       |
| `SeverityDatadogDebug`     | debug                      |