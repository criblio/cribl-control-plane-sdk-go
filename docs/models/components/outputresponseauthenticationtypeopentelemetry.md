# OutputResponseAuthenticationTypeOpenTelemetry

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.OutputResponseAuthenticationTypeOpenTelemetryNone

// Open enum: custom values can be created with a direct type cast
custom := components.OutputResponseAuthenticationTypeOpenTelemetry("custom_value")
```


## Values

| Name                                                             | Value                                                            |
| ---------------------------------------------------------------- | ---------------------------------------------------------------- |
| `OutputResponseAuthenticationTypeOpenTelemetryNone`              | none                                                             |
| `OutputResponseAuthenticationTypeOpenTelemetryBasic`             | basic                                                            |
| `OutputResponseAuthenticationTypeOpenTelemetryCredentialsSecret` | credentialsSecret                                                |
| `OutputResponseAuthenticationTypeOpenTelemetryToken`             | token                                                            |
| `OutputResponseAuthenticationTypeOpenTelemetryTextSecret`        | textSecret                                                       |
| `OutputResponseAuthenticationTypeOpenTelemetryOauthSecret`       | oauthSecret                                                      |