# RedisDeploymentTypeClusterDeploymentType

How the Redis server is configured. Defaults to Standalone

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.RedisDeploymentTypeClusterDeploymentTypeStandalone

// Open enum: custom values can be created with a direct type cast
custom := components.RedisDeploymentTypeClusterDeploymentType("custom_value")
```


## Values

| Name                                                 | Value                                                |
| ---------------------------------------------------- | ---------------------------------------------------- |
| `RedisDeploymentTypeClusterDeploymentTypeStandalone` | standalone                                           |
| `RedisDeploymentTypeClusterDeploymentTypeCluster`    | cluster                                              |
| `RedisDeploymentTypeClusterDeploymentTypeSentinel`   | sentinel                                             |