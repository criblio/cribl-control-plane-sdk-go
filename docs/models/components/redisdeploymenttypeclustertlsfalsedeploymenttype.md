# RedisDeploymentTypeClusterTLSFalseDeploymentType

How the Redis server is configured. Defaults to Standalone

## Example Usage

```go
import (
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

value := components.RedisDeploymentTypeClusterTLSFalseDeploymentTypeStandalone

// Open enum: custom values can be created with a direct type cast
custom := components.RedisDeploymentTypeClusterTLSFalseDeploymentType("custom_value")
```


## Values

| Name                                                         | Value                                                        |
| ------------------------------------------------------------ | ------------------------------------------------------------ |
| `RedisDeploymentTypeClusterTLSFalseDeploymentTypeStandalone` | standalone                                                   |
| `RedisDeploymentTypeClusterTLSFalseDeploymentTypeCluster`    | cluster                                                      |
| `RedisDeploymentTypeClusterTLSFalseDeploymentTypeSentinel`   | sentinel                                                     |