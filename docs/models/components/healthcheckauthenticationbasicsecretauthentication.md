# HealthCheckAuthenticationBasicSecretAuthentication

Authentication method for Discover and Collect REST calls. You can specify API Key–based authentication by adding the appropriate Collect headers.

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.HealthCheckAuthenticationBasicSecretAuthenticationNone

// Open enum: custom values can be created with a direct type cast
custom := components.HealthCheckAuthenticationBasicSecretAuthentication("custom_value")
```


## Values

| Name                                                            | Value                                                           |
| --------------------------------------------------------------- | --------------------------------------------------------------- |
| `HealthCheckAuthenticationBasicSecretAuthenticationNone`        | none                                                            |
| `HealthCheckAuthenticationBasicSecretAuthenticationBasic`       | basic                                                           |
| `HealthCheckAuthenticationBasicSecretAuthenticationBasicSecret` | basicSecret                                                     |
| `HealthCheckAuthenticationBasicSecretAuthenticationLogin`       | login                                                           |
| `HealthCheckAuthenticationBasicSecretAuthenticationLoginSecret` | loginSecret                                                     |
| `HealthCheckAuthenticationBasicSecretAuthenticationOauth`       | oauth                                                           |
| `HealthCheckAuthenticationBasicSecretAuthenticationOauthSecret` | oauthSecret                                                     |