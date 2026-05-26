# InputResponseAuthenticationTypeOpenTelemetry

OpenTelemetry authentication type

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.InputResponseAuthenticationTypeOpenTelemetryNone

// Open enum: custom values can be created with a direct type cast
custom := components.InputResponseAuthenticationTypeOpenTelemetry("custom_value")
```


## Values

| Name                                                            | Value                                                           |
| --------------------------------------------------------------- | --------------------------------------------------------------- |
| `InputResponseAuthenticationTypeOpenTelemetryNone`              | none                                                            |
| `InputResponseAuthenticationTypeOpenTelemetryBasic`             | basic                                                           |
| `InputResponseAuthenticationTypeOpenTelemetryCredentialsSecret` | credentialsSecret                                               |
| `InputResponseAuthenticationTypeOpenTelemetryToken`             | token                                                           |
| `InputResponseAuthenticationTypeOpenTelemetryTextSecret`        | textSecret                                                      |