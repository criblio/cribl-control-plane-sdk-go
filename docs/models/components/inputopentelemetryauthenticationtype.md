# InputOpenTelemetryAuthenticationType

OpenTelemetry authentication type

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.InputOpenTelemetryAuthenticationTypeNone

// Open enum: custom values can be created with a direct type cast
custom := components.InputOpenTelemetryAuthenticationType("custom_value")
```


## Values

| Name                                                    | Value                                                   |
| ------------------------------------------------------- | ------------------------------------------------------- |
| `InputOpenTelemetryAuthenticationTypeNone`              | none                                                    |
| `InputOpenTelemetryAuthenticationTypeBasic`             | basic                                                   |
| `InputOpenTelemetryAuthenticationTypeCredentialsSecret` | credentialsSecret                                       |
| `InputOpenTelemetryAuthenticationTypeToken`             | token                                                   |
| `InputOpenTelemetryAuthenticationTypeTextSecret`        | textSecret                                              |