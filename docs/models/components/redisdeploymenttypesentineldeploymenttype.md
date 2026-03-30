# RedisDeploymentTypeSentinelDeploymentType

How the Redis server is configured. Defaults to Standalone

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.RedisDeploymentTypeSentinelDeploymentTypeStandalone

// Open enum: custom values can be created with a direct type cast
custom := components.RedisDeploymentTypeSentinelDeploymentType("custom_value")
```


## Values

| Name                                                  | Value                                                 |
| ----------------------------------------------------- | ----------------------------------------------------- |
| `RedisDeploymentTypeSentinelDeploymentTypeStandalone` | standalone                                            |
| `RedisDeploymentTypeSentinelDeploymentTypeCluster`    | cluster                                               |
| `RedisDeploymentTypeSentinelDeploymentTypeSentinel`   | sentinel                                              |