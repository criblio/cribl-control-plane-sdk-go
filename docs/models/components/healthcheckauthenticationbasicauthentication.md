# HealthCheckAuthenticationBasicAuthentication

Authentication method for Discover and Collect REST calls. You can specify API Key–based authentication by adding the appropriate Collect headers.

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.HealthCheckAuthenticationBasicAuthenticationNone

// Open enum: custom values can be created with a direct type cast
custom := components.HealthCheckAuthenticationBasicAuthentication("custom_value")
```


## Values

| Name                                                      | Value                                                     |
| --------------------------------------------------------- | --------------------------------------------------------- |
| `HealthCheckAuthenticationBasicAuthenticationNone`        | none                                                      |
| `HealthCheckAuthenticationBasicAuthenticationBasic`       | basic                                                     |
| `HealthCheckAuthenticationBasicAuthenticationBasicSecret` | basicSecret                                               |
| `HealthCheckAuthenticationBasicAuthenticationLogin`       | login                                                     |
| `HealthCheckAuthenticationBasicAuthenticationLoginSecret` | loginSecret                                               |
| `HealthCheckAuthenticationBasicAuthenticationOauth`       | oauth                                                     |
| `HealthCheckAuthenticationBasicAuthenticationOauthSecret` | oauthSecret                                               |