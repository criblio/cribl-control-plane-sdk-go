# CreateOutputSystemByPackAuthenticationTypeOpenTelemetry

Authentication type

## Example Usage

```go
import (
	"github.com/Cribl-Community/cribl-control-plane-sdk-go/models/operations"
)

value := operations.CreateOutputSystemByPackAuthenticationTypeOpenTelemetryNone

// Open enum: custom values can be created with a direct type cast
custom := operations.CreateOutputSystemByPackAuthenticationTypeOpenTelemetry("custom_value")
```


## Values

| Name                                                                       | Value                                                                      |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `CreateOutputSystemByPackAuthenticationTypeOpenTelemetryNone`              | none                                                                       |
| `CreateOutputSystemByPackAuthenticationTypeOpenTelemetryBasic`             | basic                                                                      |
| `CreateOutputSystemByPackAuthenticationTypeOpenTelemetryCredentialsSecret` | credentialsSecret                                                          |
| `CreateOutputSystemByPackAuthenticationTypeOpenTelemetryToken`             | token                                                                      |
| `CreateOutputSystemByPackAuthenticationTypeOpenTelemetryTextSecret`        | textSecret                                                                 |
| `CreateOutputSystemByPackAuthenticationTypeOpenTelemetryOauthSecret`       | oauthSecret                                                                |