# OutputResponseOutputOpenTelemetryAuthenticationType

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.OutputResponseOutputOpenTelemetryAuthenticationTypeNone

// Open enum: custom values can be created with a direct type cast
custom := components.OutputResponseOutputOpenTelemetryAuthenticationType("custom_value")
```


## Values

| Name                                                                   | Value                                                                  |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `OutputResponseOutputOpenTelemetryAuthenticationTypeNone`              | none                                                                   |
| `OutputResponseOutputOpenTelemetryAuthenticationTypeBasic`             | basic                                                                  |
| `OutputResponseOutputOpenTelemetryAuthenticationTypeCredentialsSecret` | credentialsSecret                                                      |
| `OutputResponseOutputOpenTelemetryAuthenticationTypeToken`             | token                                                                  |
| `OutputResponseOutputOpenTelemetryAuthenticationTypeTextSecret`        | textSecret                                                             |
| `OutputResponseOutputOpenTelemetryAuthenticationTypeOauthSecret`       | oauthSecret                                                            |