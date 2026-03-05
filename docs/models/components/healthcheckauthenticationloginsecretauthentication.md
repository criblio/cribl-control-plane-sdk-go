# HealthCheckAuthenticationLoginSecretAuthentication

Authentication method for Discover and Collect REST calls. You can specify API Key–based authentication by adding the appropriate Collect headers.

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.HealthCheckAuthenticationLoginSecretAuthenticationNone

// Open enum: custom values can be created with a direct type cast
custom := components.HealthCheckAuthenticationLoginSecretAuthentication("custom_value")
```


## Values

| Name                                                            | Value                                                           |
| --------------------------------------------------------------- | --------------------------------------------------------------- |
| `HealthCheckAuthenticationLoginSecretAuthenticationNone`        | none                                                            |
| `HealthCheckAuthenticationLoginSecretAuthenticationBasic`       | basic                                                           |
| `HealthCheckAuthenticationLoginSecretAuthenticationBasicSecret` | basicSecret                                                     |
| `HealthCheckAuthenticationLoginSecretAuthenticationLogin`       | login                                                           |
| `HealthCheckAuthenticationLoginSecretAuthenticationLoginSecret` | loginSecret                                                     |
| `HealthCheckAuthenticationLoginSecretAuthenticationOauth`       | oauth                                                           |
| `HealthCheckAuthenticationLoginSecretAuthenticationOauthSecret` | oauthSecret                                                     |