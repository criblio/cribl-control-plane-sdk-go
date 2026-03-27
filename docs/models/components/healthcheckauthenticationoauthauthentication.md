# HealthCheckAuthenticationOauthAuthentication

Authentication method for Discover and Collect REST calls. You can specify API Key–based authentication by adding the appropriate Collect headers.

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.HealthCheckAuthenticationOauthAuthenticationNone

// Open enum: custom values can be created with a direct type cast
custom := components.HealthCheckAuthenticationOauthAuthentication("custom_value")
```


## Values

| Name                                                      | Value                                                     |
| --------------------------------------------------------- | --------------------------------------------------------- |
| `HealthCheckAuthenticationOauthAuthenticationNone`        | none                                                      |
| `HealthCheckAuthenticationOauthAuthenticationBasic`       | basic                                                     |
| `HealthCheckAuthenticationOauthAuthenticationBasicSecret` | basicSecret                                               |
| `HealthCheckAuthenticationOauthAuthenticationLogin`       | login                                                     |
| `HealthCheckAuthenticationOauthAuthenticationLoginSecret` | loginSecret                                               |
| `HealthCheckAuthenticationOauthAuthenticationOauth`       | oauth                                                     |
| `HealthCheckAuthenticationOauthAuthenticationOauthSecret` | oauthSecret                                               |