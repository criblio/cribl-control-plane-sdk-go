# CreateInputAuthenticationTypeOpenTelemetry

OpenTelemetry authentication type

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/operations"
)

value := operations.CreateInputAuthenticationTypeOpenTelemetryNone

// Open enum: custom values can be created with a direct type cast
custom := operations.CreateInputAuthenticationTypeOpenTelemetry("custom_value")
```


## Values

| Name                                                          | Value                                                         |
| ------------------------------------------------------------- | ------------------------------------------------------------- |
| `CreateInputAuthenticationTypeOpenTelemetryNone`              | none                                                          |
| `CreateInputAuthenticationTypeOpenTelemetryBasic`             | basic                                                         |
| `CreateInputAuthenticationTypeOpenTelemetryCredentialsSecret` | credentialsSecret                                             |
| `CreateInputAuthenticationTypeOpenTelemetryToken`             | token                                                         |
| `CreateInputAuthenticationTypeOpenTelemetryTextSecret`        | textSecret                                                    |