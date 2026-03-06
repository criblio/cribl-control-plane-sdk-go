# RedisAuthTypeTextSecretDeploymentType

How the Redis server is configured. Defaults to Standalone

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.RedisAuthTypeTextSecretDeploymentTypeStandalone

// Open enum: custom values can be created with a direct type cast
custom := components.RedisAuthTypeTextSecretDeploymentType("custom_value")
```


## Values

| Name                                              | Value                                             |
| ------------------------------------------------- | ------------------------------------------------- |
| `RedisAuthTypeTextSecretDeploymentTypeStandalone` | standalone                                        |
| `RedisAuthTypeTextSecretDeploymentTypeCluster`    | cluster                                           |
| `RedisAuthTypeTextSecretDeploymentTypeSentinel`   | sentinel                                          |