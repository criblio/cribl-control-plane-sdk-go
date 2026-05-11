# CreateOutputAuthenticationTypeOpenTelemetry

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/operations"
)

value := operations.CreateOutputAuthenticationTypeOpenTelemetryNone

// Open enum: custom values can be created with a direct type cast
custom := operations.CreateOutputAuthenticationTypeOpenTelemetry("custom_value")
```


## Values

| Name                                                           | Value                                                          |
| -------------------------------------------------------------- | -------------------------------------------------------------- |
| `CreateOutputAuthenticationTypeOpenTelemetryNone`              | none                                                           |
| `CreateOutputAuthenticationTypeOpenTelemetryBasic`             | basic                                                          |
| `CreateOutputAuthenticationTypeOpenTelemetryCredentialsSecret` | credentialsSecret                                              |
| `CreateOutputAuthenticationTypeOpenTelemetryToken`             | token                                                          |
| `CreateOutputAuthenticationTypeOpenTelemetryTextSecret`        | textSecret                                                     |
| `CreateOutputAuthenticationTypeOpenTelemetryOauthSecret`       | oauthSecret                                                    |