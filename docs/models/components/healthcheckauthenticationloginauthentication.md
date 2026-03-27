# HealthCheckAuthenticationLoginAuthentication

Authentication method for Discover and Collect REST calls. You can specify API Key–based authentication by adding the appropriate Collect headers.

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.HealthCheckAuthenticationLoginAuthenticationNone

// Open enum: custom values can be created with a direct type cast
custom := components.HealthCheckAuthenticationLoginAuthentication("custom_value")
```


## Values

| Name                                                      | Value                                                     |
| --------------------------------------------------------- | --------------------------------------------------------- |
| `HealthCheckAuthenticationLoginAuthenticationNone`        | none                                                      |
| `HealthCheckAuthenticationLoginAuthenticationBasic`       | basic                                                     |
| `HealthCheckAuthenticationLoginAuthenticationBasicSecret` | basicSecret                                               |
| `HealthCheckAuthenticationLoginAuthenticationLogin`       | login                                                     |
| `HealthCheckAuthenticationLoginAuthenticationLoginSecret` | loginSecret                                               |
| `HealthCheckAuthenticationLoginAuthenticationOauth`       | oauth                                                     |
| `HealthCheckAuthenticationLoginAuthenticationOauthSecret` | oauthSecret                                               |