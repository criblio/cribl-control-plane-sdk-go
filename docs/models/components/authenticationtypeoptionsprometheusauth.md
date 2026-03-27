# AuthenticationTypeOptionsPrometheusAuth

Remote Write authentication type

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.AuthenticationTypeOptionsPrometheusAuthNone

// Open enum: custom values can be created with a direct type cast
custom := components.AuthenticationTypeOptionsPrometheusAuth("custom_value")
```


## Values

| Name                                                       | Value                                                      |
| ---------------------------------------------------------- | ---------------------------------------------------------- |
| `AuthenticationTypeOptionsPrometheusAuthNone`              | none                                                       |
| `AuthenticationTypeOptionsPrometheusAuthBasic`             | basic                                                      |
| `AuthenticationTypeOptionsPrometheusAuthCredentialsSecret` | credentialsSecret                                          |
| `AuthenticationTypeOptionsPrometheusAuthToken`             | token                                                      |
| `AuthenticationTypeOptionsPrometheusAuthTextSecret`        | textSecret                                                 |