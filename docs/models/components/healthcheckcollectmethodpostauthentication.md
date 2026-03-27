# HealthCheckCollectMethodPostAuthentication

Authentication method for Discover and Collect REST calls. You can specify API Key–based authentication by adding the appropriate Collect headers.

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.HealthCheckCollectMethodPostAuthenticationNone

// Open enum: custom values can be created with a direct type cast
custom := components.HealthCheckCollectMethodPostAuthentication("custom_value")
```


## Values

| Name                                                    | Value                                                   |
| ------------------------------------------------------- | ------------------------------------------------------- |
| `HealthCheckCollectMethodPostAuthenticationNone`        | none                                                    |
| `HealthCheckCollectMethodPostAuthenticationBasic`       | basic                                                   |
| `HealthCheckCollectMethodPostAuthenticationBasicSecret` | basicSecret                                             |
| `HealthCheckCollectMethodPostAuthenticationLogin`       | login                                                   |
| `HealthCheckCollectMethodPostAuthenticationLoginSecret` | loginSecret                                             |
| `HealthCheckCollectMethodPostAuthenticationOauth`       | oauth                                                   |
| `HealthCheckCollectMethodPostAuthenticationOauthSecret` | oauthSecret                                             |