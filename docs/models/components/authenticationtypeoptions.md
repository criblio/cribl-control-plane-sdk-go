# AuthenticationTypeOptions

OpenTelemetry authentication type

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.AuthenticationTypeOptionsNone

// Open enum: custom values can be created with a direct type cast
custom := components.AuthenticationTypeOptions("custom_value")
```


## Values

| Name                                         | Value                                        |
| -------------------------------------------- | -------------------------------------------- |
| `AuthenticationTypeOptionsNone`              | none                                         |
| `AuthenticationTypeOptionsBasic`             | basic                                        |
| `AuthenticationTypeOptionsCredentialsSecret` | credentialsSecret                            |
| `AuthenticationTypeOptionsToken`             | token                                        |
| `AuthenticationTypeOptionsTextSecret`        | textSecret                                   |