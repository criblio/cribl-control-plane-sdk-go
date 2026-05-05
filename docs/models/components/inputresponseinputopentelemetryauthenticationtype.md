# InputResponseInputOpenTelemetryAuthenticationType

OpenTelemetry authentication type

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.InputResponseInputOpenTelemetryAuthenticationTypeNone

// Open enum: custom values can be created with a direct type cast
custom := components.InputResponseInputOpenTelemetryAuthenticationType("custom_value")
```


## Values

| Name                                                                 | Value                                                                |
| -------------------------------------------------------------------- | -------------------------------------------------------------------- |
| `InputResponseInputOpenTelemetryAuthenticationTypeNone`              | none                                                                 |
| `InputResponseInputOpenTelemetryAuthenticationTypeBasic`             | basic                                                                |
| `InputResponseInputOpenTelemetryAuthenticationTypeCredentialsSecret` | credentialsSecret                                                    |
| `InputResponseInputOpenTelemetryAuthenticationTypeToken`             | token                                                                |
| `InputResponseInputOpenTelemetryAuthenticationTypeTextSecret`        | textSecret                                                           |