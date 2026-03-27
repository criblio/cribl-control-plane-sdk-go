# HealthCheckAuthenticationOauthSecretAuthentication

Authentication method for Discover and Collect REST calls. You can specify API Key–based authentication by adding the appropriate Collect headers.

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.HealthCheckAuthenticationOauthSecretAuthenticationNone

// Open enum: custom values can be created with a direct type cast
custom := components.HealthCheckAuthenticationOauthSecretAuthentication("custom_value")
```


## Values

| Name                                                            | Value                                                           |
| --------------------------------------------------------------- | --------------------------------------------------------------- |
| `HealthCheckAuthenticationOauthSecretAuthenticationNone`        | none                                                            |
| `HealthCheckAuthenticationOauthSecretAuthenticationBasic`       | basic                                                           |
| `HealthCheckAuthenticationOauthSecretAuthenticationBasicSecret` | basicSecret                                                     |
| `HealthCheckAuthenticationOauthSecretAuthenticationLogin`       | login                                                           |
| `HealthCheckAuthenticationOauthSecretAuthenticationLoginSecret` | loginSecret                                                     |
| `HealthCheckAuthenticationOauthSecretAuthenticationOauth`       | oauth                                                           |
| `HealthCheckAuthenticationOauthSecretAuthenticationOauthSecret` | oauthSecret                                                     |