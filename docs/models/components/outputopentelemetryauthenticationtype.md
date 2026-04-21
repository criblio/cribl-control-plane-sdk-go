# OutputOpenTelemetryAuthenticationType

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.OutputOpenTelemetryAuthenticationTypeNone

// Open enum: custom values can be created with a direct type cast
custom := components.OutputOpenTelemetryAuthenticationType("custom_value")
```


## Values

| Name                                                     | Value                                                    |
| -------------------------------------------------------- | -------------------------------------------------------- |
| `OutputOpenTelemetryAuthenticationTypeNone`              | none                                                     |
| `OutputOpenTelemetryAuthenticationTypeBasic`             | basic                                                    |
| `OutputOpenTelemetryAuthenticationTypeCredentialsSecret` | credentialsSecret                                        |
| `OutputOpenTelemetryAuthenticationTypeToken`             | token                                                    |
| `OutputOpenTelemetryAuthenticationTypeTextSecret`        | textSecret                                               |
| `OutputOpenTelemetryAuthenticationTypeOauthSecret`       | oauthSecret                                              |