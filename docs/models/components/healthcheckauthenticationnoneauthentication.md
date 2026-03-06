# HealthCheckAuthenticationNoneAuthentication

Authentication method for Discover and Collect REST calls. You can specify API Key–based authentication by adding the appropriate Collect headers.

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.HealthCheckAuthenticationNoneAuthenticationNone

// Open enum: custom values can be created with a direct type cast
custom := components.HealthCheckAuthenticationNoneAuthentication("custom_value")
```


## Values

| Name                                                     | Value                                                    |
| -------------------------------------------------------- | -------------------------------------------------------- |
| `HealthCheckAuthenticationNoneAuthenticationNone`        | none                                                     |
| `HealthCheckAuthenticationNoneAuthenticationBasic`       | basic                                                    |
| `HealthCheckAuthenticationNoneAuthenticationBasicSecret` | basicSecret                                              |
| `HealthCheckAuthenticationNoneAuthenticationLogin`       | login                                                    |
| `HealthCheckAuthenticationNoneAuthenticationLoginSecret` | loginSecret                                              |
| `HealthCheckAuthenticationNoneAuthenticationOauth`       | oauth                                                    |
| `HealthCheckAuthenticationNoneAuthenticationOauthSecret` | oauthSecret                                              |