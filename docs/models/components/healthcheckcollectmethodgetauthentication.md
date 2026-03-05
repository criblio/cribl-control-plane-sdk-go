# HealthCheckCollectMethodGetAuthentication

Authentication method for Discover and Collect REST calls. You can specify API Key–based authentication by adding the appropriate Collect headers.

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.HealthCheckCollectMethodGetAuthenticationNone

// Open enum: custom values can be created with a direct type cast
custom := components.HealthCheckCollectMethodGetAuthentication("custom_value")
```


## Values

| Name                                                   | Value                                                  |
| ------------------------------------------------------ | ------------------------------------------------------ |
| `HealthCheckCollectMethodGetAuthenticationNone`        | none                                                   |
| `HealthCheckCollectMethodGetAuthenticationBasic`       | basic                                                  |
| `HealthCheckCollectMethodGetAuthenticationBasicSecret` | basicSecret                                            |
| `HealthCheckCollectMethodGetAuthenticationLogin`       | login                                                  |
| `HealthCheckCollectMethodGetAuthenticationLoginSecret` | loginSecret                                            |
| `HealthCheckCollectMethodGetAuthenticationOauth`       | oauth                                                  |
| `HealthCheckCollectMethodGetAuthenticationOauthSecret` | oauthSecret                                            |