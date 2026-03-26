# RedisDeploymentTypeClusterTLSTrueDeploymentType

How the Redis server is configured. Defaults to Standalone

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.RedisDeploymentTypeClusterTLSTrueDeploymentTypeStandalone

// Open enum: custom values can be created with a direct type cast
custom := components.RedisDeploymentTypeClusterTLSTrueDeploymentType("custom_value")
```


## Values

| Name                                                        | Value                                                       |
| ----------------------------------------------------------- | ----------------------------------------------------------- |
| `RedisDeploymentTypeClusterTLSTrueDeploymentTypeStandalone` | standalone                                                  |
| `RedisDeploymentTypeClusterTLSTrueDeploymentTypeCluster`    | cluster                                                     |
| `RedisDeploymentTypeClusterTLSTrueDeploymentTypeSentinel`   | sentinel                                                    |