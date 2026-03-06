# HealthCheckCollectMethodPostWithBodyAuthentication

Authentication method for Discover and Collect REST calls. You can specify API Key–based authentication by adding the appropriate Collect headers.

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.HealthCheckCollectMethodPostWithBodyAuthenticationNone

// Open enum: custom values can be created with a direct type cast
custom := components.HealthCheckCollectMethodPostWithBodyAuthentication("custom_value")
```


## Values

| Name                                                            | Value                                                           |
| --------------------------------------------------------------- | --------------------------------------------------------------- |
| `HealthCheckCollectMethodPostWithBodyAuthenticationNone`        | none                                                            |
| `HealthCheckCollectMethodPostWithBodyAuthenticationBasic`       | basic                                                           |
| `HealthCheckCollectMethodPostWithBodyAuthenticationBasicSecret` | basicSecret                                                     |
| `HealthCheckCollectMethodPostWithBodyAuthenticationLogin`       | login                                                           |
| `HealthCheckCollectMethodPostWithBodyAuthenticationLoginSecret` | loginSecret                                                     |
| `HealthCheckCollectMethodPostWithBodyAuthenticationOauth`       | oauth                                                           |
| `HealthCheckCollectMethodPostWithBodyAuthenticationOauthSecret` | oauthSecret                                                     |