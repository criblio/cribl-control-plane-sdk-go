# RedisDeploymentTypeClusterAuthenticationMethod

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.RedisDeploymentTypeClusterAuthenticationMethodNone

// Open enum: custom values can be created with a direct type cast
custom := components.RedisDeploymentTypeClusterAuthenticationMethod("custom_value")
```


## Values

| Name                                                              | Value                                                             |
| ----------------------------------------------------------------- | ----------------------------------------------------------------- |
| `RedisDeploymentTypeClusterAuthenticationMethodNone`              | none                                                              |
| `RedisDeploymentTypeClusterAuthenticationMethodManual`            | manual                                                            |
| `RedisDeploymentTypeClusterAuthenticationMethodCredentialsSecret` | credentialsSecret                                                 |
| `RedisDeploymentTypeClusterAuthenticationMethodTextSecret`        | textSecret                                                        |