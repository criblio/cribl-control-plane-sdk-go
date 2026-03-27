# AuthenticationTypeOptionsPrometheusAuth1

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.AuthenticationTypeOptionsPrometheusAuth1None

// Open enum: custom values can be created with a direct type cast
custom := components.AuthenticationTypeOptionsPrometheusAuth1("custom_value")
```


## Values

| Name                                                        | Value                                                       |
| ----------------------------------------------------------- | ----------------------------------------------------------- |
| `AuthenticationTypeOptionsPrometheusAuth1None`              | none                                                        |
| `AuthenticationTypeOptionsPrometheusAuth1Token`             | token                                                       |
| `AuthenticationTypeOptionsPrometheusAuth1TextSecret`        | textSecret                                                  |
| `AuthenticationTypeOptionsPrometheusAuth1Basic`             | basic                                                       |
| `AuthenticationTypeOptionsPrometheusAuth1CredentialsSecret` | credentialsSecret                                           |